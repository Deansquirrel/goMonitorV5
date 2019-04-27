package action

import "github.com/Deansquirrel/goMonitorV5/object"

type IAction interface {
	//操作接口
	Do(oprType object.OprType, id string) error
}
