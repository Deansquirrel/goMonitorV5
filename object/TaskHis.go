package object

import "time"

type ITaskHis interface {
	GetHisDataId() string
}

type IntHisData struct {
	FId       string
	FConfigId string
	FNum      int
	FContent  string
	FOprTime  time.Time
}

func (i *IntHisData) GetHisDataId() string {
	return i.FId
}

type CrmDzXfTestHisData struct {
	FId       string
	FConfigId string
	FUseTime  int
	FHttpCode int
	FContent  string
	FOprTime  time.Time
}

type WebStateHisData struct {
	FId       string
	FConfigId string
	FUseTime  int
	FHttpCode int
	FContent  string
	FOprTime  time.Time
}
