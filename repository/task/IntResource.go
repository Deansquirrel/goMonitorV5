package task

import (
	"database/sql"
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/object"
)

import log "github.com/Deansquirrel/goToolLog"

const sqlGetIntTaskConfig = "" +
	"SELECT B.[FId],B.[FServer],B.[FPort],B.[FDbName],B.[FDbUser]," +
	"B.[FDbPwd],B.[FSearch],B.[FCron],B.[FCheckMax],B.[FCheckMin]," +
	"B.[FMsgTitle],B.[FMsgContent]" +
	" FROM [MConfig] A" +
	" INNER JOIN [IntTaskConfig] B ON A.[FId] = B.[FId]"

const sqlGetIntTaskConfigById = "" +
	"SELECT B.[FId],B.[FServer],B.[FPort],B.[FDbName],B.[FDbUser]," +
	"B.[FDbPwd],B.[FSearch],B.[FCron],B.[FCheckMax],B.[FCheckMin]," +
	"B.[FMsgTitle],B.[FMsgContent]" +
	" FROM [MConfig] A" +
	" INNER JOIN [IntTaskConfig] B ON A.[FId] = B.[FId]" +
	" WHERE B.[FId]=?"

type intResource struct {
}

func (r *intResource) GetSqlGetConfig() string {
	return sqlGetIntTaskConfigById
}

func (r *intResource) GetSqlGetConfigList() string {
	return sqlGetIntTaskConfig
}

func (r *intResource) DataWrapper(rows *sql.Rows) ([]object.ITaskConfig, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fServer, fDbName, fDbUser, fDbPwd, fSearch, fCron, fMsgTitle, fMsgContent string
	var fPort, fCheckMax, fCheckMin int
	resultList := make([]object.ITaskConfig, 0)
	var err error
	for rows.Next() {
		err = rows.Scan(
			&fId, &fServer, &fPort, &fDbName, &fDbUser,
			&fDbPwd, &fSearch, &fCron, &fCheckMax, &fCheckMin,
			&fMsgTitle, &fMsgContent)
		if err != nil {
			log.Error(fmt.Sprintf("convert data error: %s", err.Error()))
			return nil, err
		}
		config := object.IntTaskConfig{
			FId:         fId,
			FServer:     fServer,
			FPort:       fPort,
			FDbName:     fDbName,
			FDbUser:     fDbUser,
			FDbPwd:      fDbPwd,
			FSearch:     fSearch,
			FCron:       fCron,
			FCheckMax:   fCheckMax,
			FCheckMin:   fCheckMin,
			FMsgTitle:   fMsgTitle,
			FMsgContent: fMsgContent,
		}
		resultList = append(resultList, &config)
	}
	if rows.Err() != nil {
		log.Error(fmt.Sprintf("convert data error: %s", rows.Err().Error()))
		return nil, rows.Err()
	}
	return resultList, nil
}
