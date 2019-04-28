package taskHis

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/object"
	rep "github.com/Deansquirrel/goMonitorV5/repository"
	"github.com/Deansquirrel/goToolCommon"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

type IHisRepository interface {
	SetHis(data object.ITaskHis) error
	ClearHis(t time.Duration) error
}

func NewHisRepository(taskType object.TaskType) (IHisRepository, error) {
	switch taskType {
	case object.Int:
		return &hisRepository{&intHisResource{}}, nil
	case object.IntD:
		return nil, errors.New(fmt.Sprintf("unexpected task type %d", taskType))
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

type iHisRepositoryResource interface {
	GetSqlSetHis() string
	GetSqlClearHis() string
	GetHisSetArgs(data object.ITaskHis) ([]interface{}, error)
}

type hisRepository struct {
	resource iHisRepositoryResource
}

func (hr *hisRepository) SetHis(data object.ITaskHis) error {
	args, err := hr.resource.GetHisSetArgs(data)
	if err != nil {
		return err
	}
	comm := rep.Common{}
	err = comm.SetRowsBySQL(hr.resource.GetSqlSetHis(), args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (hr *hisRepository) ClearHis(t time.Duration) error {
	dateP := goToolCommon.GetDateTimeStr(time.Now().Add(-t))
	comm := rep.Common{}
	err := comm.SetRowsBySQL(hr.resource.GetSqlClearHis(), dateP)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
