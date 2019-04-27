package notify

import (
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/object"
	"github.com/Deansquirrel/goMonitorV5/repository"
	"github.com/Deansquirrel/goToolCommon"
	"strings"
)

import log "github.com/Deansquirrel/goToolLog"

const (
	sqlGetNotifyList = "" +
		"SELECT [DingTalkRobotId] " +
		"FROM NotifyList " +
		"WHERE [TaskId] = ? or TaskId = '-1'"
)

type Common struct {
}

//获取通知配置列表
// - id TaskId
func (c *Common) GetNotifyList(id string) (*object.NotifyConfigList, error) {
	r := repository.Common{}
	rows, err := r.GetRowsBySQL(sqlGetNotifyList, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	dingTalkRobotIdList := make([]string, 0)
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			log.Error(fmt.Sprintf("[GetNotifyList]read rows data error: %s", err.Error()))
			return nil, err
		}
		idList := strings.Split(id, "|")
		idList = goToolCommon.ClearBlock(idList)
		for _, idStr := range idList {
			dingTalkRobotIdList = append(dingTalkRobotIdList, idStr)
		}
	}
	return &object.NotifyConfigList{
		DingTalkRobotId: dingTalkRobotIdList,
	}, nil
}
