package task

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/object"
	rep "github.com/Deansquirrel/goMonitorV5/repository"
)

import log "github.com/Deansquirrel/goToolLog"

type IRepository interface {
	GetConfig(id string) (object.ITaskConfig, error)
	GetConfigList() ([]object.ITaskConfig, error)
}

func NewRepository(taskType object.TaskType) (IRepository, error) {
	switch taskType {
	case object.Int:
		return &repository{&intResource{}}, nil
	case object.IntD:
		return &repository{&intDResource{}}, nil
	case object.CrmDzXfTest:
		return nil, errors.New(fmt.Sprintf("unexpected task type %d", taskType))
	case object.Health:
		return nil, errors.New(fmt.Sprintf("unexpected task type %d", taskType))
	case object.WebState:
		return nil, errors.New(fmt.Sprintf("unexpected task type %d", taskType))
	default:
		return nil, errors.New(fmt.Sprintf("unexpected task type %d", taskType))
	}
}

type iRepositoryResource interface {
	GetSqlGetConfig() string
	GetSqlGetConfigList() string
	DataWrapper(rows *sql.Rows) ([]object.ITaskConfig, error)
}

type repository struct {
	resource iRepositoryResource
}

func (r *repository) GetConfig(id string) (object.ITaskConfig, error) {
	c := rep.Common{}
	rows, err := c.GetRowsBySQL(r.resource.GetSqlGetConfig(), id)
	if err != nil {
		log.Error(fmt.Sprintf("get config error,sql [%s],id [%s]", r.resource.GetSqlGetConfig(), id))
		return nil, err
	}
	list, err := r.resource.DataWrapper(rows)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	} else {
		return nil, nil
	}
}

func (r *repository) GetConfigList() ([]object.ITaskConfig, error) {
	c := rep.Common{}
	rows, err := c.GetRowsBySQL(r.resource.GetSqlGetConfigList())
	if err != nil {
		log.Error(fmt.Sprintf("get config list error,sql [%s]", r.resource.GetSqlGetConfig()))
		return nil, err
	}
	list, err := r.resource.DataWrapper(rows)
	if err != nil {
		return nil, err
	}
	return list, nil
}
