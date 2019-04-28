package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/notify"
	"github.com/Deansquirrel/goMonitorV5/object"
	"github.com/Deansquirrel/goMonitorV5/worker"
	"github.com/robfig/cron"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

func NewTask(taskType object.TaskType, config object.ITaskConfig) object.ITask {
	ctx, cancel := context.WithCancel(context.Background())
	t := task{
		id:       "",
		taskType: taskType,
		config:   config,
		cron:     nil,
		running:  false,
		err:      nil,

		ctx:    ctx,
		cancel: cancel,
	}
	if config != nil {
		t.id = config.GetConfigId()
	}
	t.Start()
	return &t
}

type task struct {
	id       string
	taskType object.TaskType
	config   object.ITaskConfig
	cron     *cron.Cron
	running  bool
	err      error

	ctx    context.Context
	cancel func()
}

func (t *task) GetTaskType() object.TaskType {
	return t.taskType
}

func (t *task) Start() {
	if t.running {
		return
	}
	if t.config == nil {
		t.err = errors.New(fmt.Sprintf("config is nil"))
		return
	}
	w := worker.NewWorker(t.config)
	if w == nil {
		t.err = errors.New(fmt.Sprintf("get worker return nil"))
		return
	}
	c, err := t.getCron(w)
	if err != nil {
		t.err = errors.New(fmt.Sprintf("get cron error: %s", err.Error()))
		return
	}
	t.cron = c
	t.cron.Start()

	//创建清理历史数据任务
	go func() {
		c := cron.New()
		err := c.AddFunc("0 0 * * * ?", func() {
			delErr := w.DelHisData()
			if delErr != nil {
				errMsg := fmt.Sprintf("del his task error,id %s,error %s'", t.id, err.Error())
				log.Error(errMsg)
				_, _, _, _, _ = notify.SendMsg(t.id, errMsg)
			}
		})
		if err != nil {
			errMsg := fmt.Sprintf("create del his task error,id %s,error %s'", t.id, err.Error())
			log.Error(errMsg)
			_, _, _, _, _ = notify.SendMsg(t.id, errMsg)
		} else {
			c.Start()
			select {
			case <-t.ctx.Done():
				return
			}
		}
	}()

	t.running = true
}

func (t *task) Stop() {
	t.running = false
	t.cancel()
	t.cron.Stop()
}

func (t *task) IsEqual(config object.ITaskConfig) bool {
	return t.config.IsEqual(config)
}

func (t *task) GetTaskId() string {
	return t.id
}

func (t *task) IsRunning() bool {
	return t.running
}

func (t *task) GetError() error {
	return t.err
}

func (t *task) getCron(w worker.IWorker) (*cron.Cron, error) {
	c := cron.New()
	err := c.AddFunc(t.config.GetSpec(), func() {
		msg, hisData := w.GetMsg()
		if msg != "" {
			//发送消息通知
			_, _, _, _, _ = notify.SendMsg(t.id, msg)
			//检查操作
			err := w.CheckAction()
			if err != nil {
				errMsg := fmt.Sprintf("check action error,id %s,error %s'", t.id, err.Error())
				log.Error(errMsg)
				t, s, f, e, eMap := notify.SendMsg(t.id, errMsg)
				log.Warn(fmt.Sprintf("send err result: total %d success %d fail %d err %s errMap %v",
					t, s, f, e.Error(), eMap))
			}
		}

		//保存查询数据
		if hisData != nil {
			err := w.SaveData(hisData)
			if err != nil {
				errMsg := fmt.Sprintf("save his data error,id %s,error %s'", t.id, err.Error())
				log.Error(errMsg)
				t, s, f, e, eMap := notify.SendMsg(t.id, errMsg)
				log.Warn(fmt.Sprintf("send err result: total %d success %d fail %d err %s errMap %v",
					t, s, f, e.Error(), eMap))
			}
		}

	})
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (t *task) Prev() time.Time {
	if t.cron != nil {
		eList := t.cron.Entries()
		if len(eList) > 0 {
			return eList[0].Prev
		}
	}
	return time.Date(1900, 0, 1, 0, 0, 0, 0, time.Local)
}

func (t *task) Next() time.Time {
	if t.cron != nil {
		eList := t.cron.Entries()
		if len(eList) > 0 {
			return eList[0].Next
		}
	}
	return time.Date(1900, 0, 1, 0, 0, 0, 0, time.Local)
}
