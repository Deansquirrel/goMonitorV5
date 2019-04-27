package worker

import "github.com/Deansquirrel/goMonitorV5/object"

type IWorker interface {
	Run()
}

func NewWorker(config object.ITaskConfig) IWorker {
	return nil
}
