package service

import (
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/object"
	"github.com/Deansquirrel/goMonitorV5/taskService"
	"github.com/Deansquirrel/goMonitorV5/worker"
	"github.com/robfig/cron"
)

import log "github.com/Deansquirrel/goToolLog"

type TaskServer struct {
	taskManager object.ITaskManager
}

func (t *TaskServer) Start() {
	t.taskManager = object.NewTaskManager()
	taskConfigService := taskService.TaskConfigService{}
	for _, taskType := range object.TaskTypeList {
		configList := taskConfigService.GetConfigList(taskType)
		for _, config := range configList {
			w := worker.NewWorker(config)
			c := cron.New()
			task := object.NewTask(taskType, config, c)
			err := c.AddJob(config.GetSpec(), w)
			if err != nil {
				log.Error(fmt.Sprintf("get cron error: %s", err.Error()))
				task.SetError(err)
			} else {
				task.Start()
			}
			t.taskManager.GetChRegister() <- task
		}
	}
}
