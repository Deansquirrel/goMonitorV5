package taskService

import "github.com/Deansquirrel/goMonitorV5/object"

type TaskConfigService struct {
}

func (cs *TaskConfigService) GetConfig(id string, taskType object.TaskType) object.ITaskConfig {
	//TODO
	return nil
}

func (cs *TaskConfigService) GetConfigList(taskType object.TaskType) []object.ITaskConfig {
	//TODO
	return nil
}
