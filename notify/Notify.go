package notify

import (
	"fmt"
	repNotify "github.com/Deansquirrel/goMonitorV5/repository/notify"
)

import log "github.com/Deansquirrel/goToolLog"

type notify struct {
}

//获取iNotify列表
// - id TaskId
func (n *notify) GetNotifyList(id string) ([]iNotify, error) {
	c := repNotify.Common{}
	notifyConfigList, err := c.GetNotifyList(id)
	if err != nil {
		return nil, err
	}
	rList := make([]iNotify, 0)
	repDingTalkRobot := repNotify.DingTalkRobot{}
	for _, id := range notifyConfigList.DingTalkRobotId {
		log.Debug(id)
		configData, err := repDingTalkRobot.GetConfigData(id)
		if err != nil {
			return nil, err
		}
		if configData != nil {
			rList = append(rList, &dingTalkRobot{config: configData})
		} else {
			log.Warn(fmt.Sprintf("ID 为 [%s] 的DingTalkRobot的配置不存在", id))
		}
	}
	return rList, nil
}
