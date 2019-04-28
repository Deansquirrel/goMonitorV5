package taskService

import (
	"encoding/json"
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/global"
	"github.com/Deansquirrel/goMonitorV5/object"
	"github.com/Deansquirrel/goMonitorV5/task"
)

import log "github.com/Deansquirrel/goToolLog"

type TaskService struct {
	taskManager object.ITaskManager
}

func (t *TaskService) Start() {
	log.Debug("TaskService Start")
	defer log.Debug("TaskService Start Complete")
	t.start()
}

func (t *TaskService) Stop() {
	log.Debug("TaskService Stop")
	defer log.Debug("TaskService Stop Complete")
	t.stop()
}

func (t *TaskService) start() {
	if t.taskManager != nil {
		return
	}
	t.taskManager = object.NewTaskManager()

	t.startTask(object.Int)
	t.startTask(object.CrmDzXfTest)
	t.startTask(object.Health)
	t.startTask(object.WebState)

	go func() {
		select {
		case <-global.Ctx.Done():
			t.Stop()
		}
	}()
}

func (t *TaskService) startTask(taskType object.TaskType) {
	tcs := TaskConfigService{}
	taskConfigList, err := tcs.GetConfigList(taskType)
	if err != nil {
		log.Error(fmt.Sprintf("get config list error,type: %d,error: %s", taskType, err.Error()))
		return
	}
	for _, taskConfig := range taskConfigList {
		taskObject := task.NewTask(taskType, taskConfig)
		configStr, err := json.Marshal(taskConfig)
		if err != nil {
			log.Error(fmt.Sprintf("get config str error: %s", err.Error()))
		} else {
			log.Info(fmt.Sprintf("add task %s, config: %s", taskObject.GetTaskId(), configStr))
		}
		t.addTask(taskObject)
	}
}

func (t *TaskService) stop() {
	if t.taskManager != nil {
		t.taskManager.GetChClose() <- struct{}{}
	}
	t.taskManager = nil
}

func (t *TaskService) addTask(task object.ITask) {
	t.taskManager.GetChRegister() <- task
}

func (t *TaskService) delTask(id string) {
	t.taskManager.GetChUnregister() <- id
}
