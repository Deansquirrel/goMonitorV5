package object

import (
	"github.com/robfig/cron"
	"time"
)

type task struct {
	id       string
	taskType TaskType
	config   ITaskConfig
	cron     *cron.Cron
	running  bool
	last     time.Time
	err      error
}

func NewTask(taskType TaskType, config ITaskConfig, cron *cron.Cron) *task {
	return &task{
		id:       config.GetConfigId(),
		taskType: taskType,
		config:   config,
		cron:     cron,
		running:  false,
		last:     time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local),
		err:      nil,
	}
}

func (t *task) GetTaskType() TaskType {
	return t.taskType
}

func (t *task) Start() {
	if t.cron != nil && t.running == false {
		t.cron.Start()
		t.running = true
	}
}
func (t *task) Stop() {
	if t.cron != nil && t.running == true {
		t.cron.Stop()
		t.running = false
	}
}

func (t *task) IsEqual(config ITaskConfig) bool {
	return false
}

func (t *task) GetTaskId() string {
	return t.id
}

func (t *task) IsRunning() bool {
	return t.running
}

func (t *task) SetError(err error) {
	t.err = err
}

func (t *task) GetError() error {
	return t.err
}
