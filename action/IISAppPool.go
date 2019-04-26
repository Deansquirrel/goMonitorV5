package action

import "errors"

type IISAppPoolConfig struct {
}

type iisAppPool struct {
}

//操作接口
func (action *iisAppPool) Do(oprType OprType, id string) error {
	switch oprType {
	case Restart:
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
