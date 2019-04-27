package object

type NotifyConfigList struct {
	DingTalkRobotId []string
}

type DingTalkRobotConfigData struct {
	FId         string
	FWebHookKey string
	FAtMobiles  string
	FIsAtAll    int
}

func (d *DingTalkRobotConfigData) GetNotifyId() string {
	return d.FId
}
