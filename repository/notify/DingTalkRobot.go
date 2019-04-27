package notify

import (
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/object"
	"github.com/Deansquirrel/goMonitorV5/repository"
)

import log "github.com/Deansquirrel/goToolLog"

//const SqlGetDingTalkRobot = "" +
//	"SELECT B.[FId],B.[FWebHookKey],B.[FAtMobiles],B.[FIsAtAll]" +
//	" FROM [NConfig] A" +
//	" INNER JOIN [DingTalkRobot] B ON A.[FId] = B.[FId]"

const sqlGetDingTalkRobotById = "" +
	"SELECT B.[FId],B.[FWebHookKey],B.[FAtMobiles],B.[FIsAtAll]" +
	" FROM [NConfig] A" +
	" INNER JOIN [DingTalkRobot] B ON A.[FId] = B.[FId]" +
	" WHERE A.[FId]=?"

//const SqlGetDingTalkRobotByIdList = "" +
//	"SELECT B.[FId],B.[FWebHookKey],B.[FAtMobiles],B.[FIsAtAll]" +
//	" FROM [NConfig] A" +
//	" INNER JOIN [DingTalkRobot] B ON A.[FId] = B.[FId]" +
//	" WHERE A.[FId] in (%s)"

type DingTalkRobot struct {
}

func (d *DingTalkRobot) GetConfigData(id string) (*object.DingTalkRobotConfigData, error) {
	c := repository.Common{}
	rows, err := c.GetRowsBySQL(sqlGetDingTalkRobotById, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var fId, fWebHookKey, fAtMobiles string
	var fIsAtAll int
	rList := make([]*object.DingTalkRobotConfigData, 0)
	for rows.Next() {
		err = rows.Scan(&fId, &fWebHookKey, &fAtMobiles, &fIsAtAll)
		if err != nil {
			log.Error(fmt.Sprintf("read rows data error: %s", err.Error()))
			return nil, err
		}
		config := object.DingTalkRobotConfigData{
			FId:         fId,
			FWebHookKey: fWebHookKey,
			FAtMobiles:  fAtMobiles,
			FIsAtAll:    fIsAtAll,
		}
		rList = append(rList, &config)
	}
	if rows.Err() != nil {
		log.Error(fmt.Sprintf("read rows data error: %s", err.Error()))
		return nil, rows.Err()
	}
	if len(rList) > 0 {
		return rList[0], nil
	} else {
		return nil, nil
	}
}
