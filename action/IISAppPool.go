package action

import (
	"errors"
	"github.com/Deansquirrel/goMonitorV5/object"
)

type IISAppPoolConfig struct {
}

type iisAppPool struct {
}

//操作接口
func (action *iisAppPool) Do(oprType object.OprType, id string) error {
	switch oprType {
	case object.Restart:
		return action.restart(id)
	default:
		return errors.New("unexpected opr type")
	}
}

//根据ID获取配置
func (action *iisAppPool) getConfig(id string) *IISAppPoolConfig {
	//TODO
	return nil
}

//重启
func (action *iisAppPool) restart(id string) error {
	//TODO
	return nil
}
