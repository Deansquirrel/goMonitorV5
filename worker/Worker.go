package worker

import (
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/object"
	"reflect"
)

import log "github.com/Deansquirrel/goToolLog"

type IWorker interface {
	GetMsg() (string, object.ITaskHis)
	SaveData(data object.ITaskHis) error
	DelHisData() error
	CheckAction() error
}

func NewWorker(config object.ITaskConfig) IWorker {
	switch reflect.TypeOf(config).String() {
	case "*object.IntTaskConfig":
		return &intWorker{config.(*object.IntTaskConfig)}
	default:
		log.Warn(fmt.Sprintf("unexpected task config type: %s", reflect.TypeOf(config).String()))
		return nil
	}
}
