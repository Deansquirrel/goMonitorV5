package taskService

import (
	"github.com/Deansquirrel/goMonitorV5/object"
	"github.com/Deansquirrel/goMonitorV5/repository/task"
)

type TaskConfigService struct {
}

func (cs *TaskConfigService) GetConfig(id string, taskType object.TaskType) (object.ITaskConfig, error) {
	rep, err := task.NewRepository(taskType)
	if err != nil {
		return nil, err
	}
	config, err := rep.GetConfig(id)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (cs *TaskConfigService) GetConfigList(taskType object.TaskType) ([]object.ITaskConfig, error) {
	rep, err := task.NewRepository(taskType)
	if err != nil {
		return nil, err
	}
	list, err := rep.GetConfigList()
	if err != nil {
		return nil, err
	}
	return list, nil
}
