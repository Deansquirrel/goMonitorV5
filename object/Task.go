package object

import (
	"time"
)

type ITask interface {
	GetTaskId() string
	GetTaskType() TaskType
	IsEqual(config ITaskConfig) bool
	IsRunning() bool
	Prev() time.Time
	Next() time.Time
	Start()
	Stop()
	GetError() error
}
