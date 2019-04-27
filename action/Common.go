package action

import (
	"errors"
	"github.com/Deansquirrel/goMonitorV5/object"
)

func NewAction(actionType object.ActionType) (IAction, error) {
	switch actionType {
	case object.IISAppPool:
		return &iisAppPool{}, nil
	case object.WindowsService:
		return &windowsService{}, nil
	default:
		return nil, errors.New("unexpected action type")
	}
}

func Opr(actionType object.ActionType, id string, oprType OprType) error {
	a, err := NewAction(actionType)
	if err != nil {
		return err
	}
	return a.Do(oprType, id)
}
