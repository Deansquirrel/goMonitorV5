package action

import "errors"

type Type int

const (
	WindowsService Type = iota
	IISAppPool
)

func NewAction(actionType Type) (IAction, error) {
	switch actionType {
	case IISAppPool:
		return &iisAppPool{}, nil
	case WindowsService:
		return &windowsService{}, nil
	default:
		return nil, errors.New("unexpected action type")
	}
}
