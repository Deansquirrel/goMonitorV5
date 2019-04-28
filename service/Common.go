package service

import (
	"github.com/Deansquirrel/goMonitorV5/service/taskService"
)

var T *taskService.TaskService

//启动服务内容
func Start() error {
	go func() {
		T = &taskService.TaskService{}
		T.Start()
	}()
	return nil
}
