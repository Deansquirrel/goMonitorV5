package object

import (
	"reflect"
)

type ITaskConfig interface {
	GetConfigId() string
	GetSpec() string
	IsEqual(c ITaskConfig) bool
}

type IntTaskConfig struct {
	FId         string
	FServer     string
	FPort       int
	FDbName     string
	FDbUser     string
	FDbPwd      string
	FSearch     string
	FCron       string
	FCheckMax   int
	FCheckMin   int
	FMsgTitle   string
	FMsgContent string
}

func (configData *IntTaskConfig) GetSpec() string {
	return configData.FCron
}

func (configData *IntTaskConfig) GetConfigId() string {
	return configData.FId
}

func (configData *IntTaskConfig) IsEqual(d ITaskConfig) bool {
	switch reflect.TypeOf(d).String() {
	case "*IntTaskConfig":
		c, ok := d.(*IntTaskConfig)
		if !ok {
			return false
		}
		if configData.FId != c.FId ||
			configData.FServer != c.FServer ||
			configData.FPort != c.FPort ||
			configData.FDbName != c.FDbName ||
			configData.FDbUser != c.FDbUser ||
			configData.FDbPwd != c.FDbPwd ||
			configData.FSearch != c.FSearch ||
			configData.FCron != c.FCron ||
			configData.FCheckMax != c.FCheckMax ||
			configData.FCheckMin != c.FCheckMin ||
			configData.FMsgTitle != c.FMsgTitle ||
			configData.FMsgContent != c.FMsgContent {
			return false
		}
		return true
	default:
		return false
	}
}

type IntDTaskConfig struct {
	FId        string
	FMsgSearch string
}

func (configData *IntDTaskConfig) GetSpec() string {
	return ""
}

func (configData *IntDTaskConfig) GetConfigId() string {
	return configData.FId
}

func (configData *IntDTaskConfig) IsEqual(d ITaskConfig) bool {
	switch reflect.TypeOf(d).String() {
	case "*object.IntDTaskConfig":
		c, ok := d.(*IntDTaskConfig)
		if !ok {
			return false
		}
		if configData.FId != c.FId ||
			configData.FMsgSearch != c.FMsgSearch {
			return false
		}
		return true
	default:
		return false
	}
}
