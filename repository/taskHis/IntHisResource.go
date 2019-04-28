package taskHis

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/object"
	"reflect"
)

const sqlSetIntTaskHis = "" +
	"INSERT INTO [IntTaskHis]([FId],[FConfigId],[FNum],[FContent])" +
	" SELECT ?,?,?,?"

const sqlDelIntTaskHisByOprTime = "" +
	"DELETE FROM [IntTaskHis]" +
	" WHERE [FOprTime] < ?"

type intHisResource struct {
}

func (intHisResource) GetSqlSetHis() string {
	return sqlSetIntTaskHis
}

func (intHisResource) GetSqlClearHis() string {
	return sqlDelIntTaskHisByOprTime
}

func (intHisResource) GetHisSetArgs(data object.ITaskHis) ([]interface{}, error) {
	switch reflect.TypeOf(data).String() {
	case "*object.IntHisData":
		iHisData, ok := data.(*object.IntHisData)
		if ok {
			if len(iHisData.FContent) > 4000 {
				iHisData.FContent = iHisData.FContent[:4000]
			}
			result := make([]interface{}, 0)
			result = append(result, iHisData.FId)
			result = append(result, iHisData.FConfigId)
			result = append(result, iHisData.FNum)
			result = append(result, iHisData.FContent)
			return result, nil
		} else {
			return nil, errors.New("convert IntHisData error")
		}
	default:
		return nil, errors.New(fmt.Sprintf("type error,expected %s,get %s",
			"*object.IntHisData",
			reflect.TypeOf(data).String()))
	}
}
