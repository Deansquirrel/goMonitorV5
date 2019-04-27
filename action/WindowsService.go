package action

import (
	"errors"
	"github.com/Deansquirrel/goMonitorV5/object"
)

type WindowsServiceConfig struct {
}

type windowsService struct {
}

//操作接口
func (action *windowsService) Do(oprType object.OprType, id string) error {
	switch oprType {
	case object.Restart:
		return action.restart(id)
	default:
		return errors.New("unexpected opr type")
	}
}

//根据ID获取配置
func (action *windowsService) getConfig(id string) *WindowsServiceConfig {
	//TODO
	return nil
}

//重启
func (action *windowsService) restart(id string) error {
	//TODO
	return nil
}
