package task

import (
	"database/sql"
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/object"
)

import log "github.com/Deansquirrel/goToolLog"

const sqlGetIntTaskDConfig = "" +
	"SELECT [FID],[FMsgSearch] " +
	"FROM [IntTaskDConfig]"

const sqlGetIntTaskDConfigById = "" +
	"SELECT [FID],[FMsgSearch] " +
	"FROM [IntTaskDConfig] " +
	"WHERE [FId]=?"

type intDResource struct {
}

func (intDResource) GetSqlGetConfig() string {
	return sqlGetIntTaskDConfigById
}

func (intDResource) GetSqlGetConfigList() string {
	return sqlGetIntTaskDConfig
}

func (intDResource) DataWrapper(rows *sql.Rows) ([]object.ITaskConfig, error) {
	defer func() {
		_ = rows.Close()
	}()
	var fId, fMsgSearch sql.NullString
	resultList := make([]object.ITaskConfig, 0)
	for rows.Next() {
		err := rows.Scan(&fId, &fMsgSearch)
		if err != nil {
			log.Error(fmt.Sprintf("convert data error: %s", err.Error()))
			return nil, err
		}
		config := object.IntDTaskConfig{}
		config.FId = "Null"
		if fId.Valid {
			config.FId = fId.String
		}
		config.FMsgSearch = "Null"
		if fMsgSearch.Valid {
			config.FMsgSearch = fMsgSearch.String
		}
		resultList = append(resultList, &config)
	}
	if rows.Err() != nil {
		log.Error(fmt.Sprintf("convert data error: %s", rows.Err().Error()))
		return nil, rows.Err()
	}
	return resultList, nil
}
