package object

type ITask interface {
	GetTaskId() string
	GetTaskType() TaskType
	IsEqual(config ITaskConfig) bool
	IsRunning() bool
	Start()
	Stop()
	SetError(err error)
	GetError() error
}
