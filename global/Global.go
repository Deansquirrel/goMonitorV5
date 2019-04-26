package global

import (
	"context"
	"github.com/Deansquirrel/goMonitorV5/object"
)

//版本信息
const (
	//PreVersion = "0.0.0 Build20190101"
	//TestVersion = "0.0.0 Build20190101"
	Version = "0.0.0 Build20190101"
)

const (
	//http连接超时时间（秒）
	HttpConnectTimeout = 30
)

var Ctx context.Context
var Cancel func()

//程序启动参数
var Args *object.ProgramArgs

//配置文件是否存在
//var IsConfigExists bool
//系统参数
var SysConfig *object.SystemConfig
