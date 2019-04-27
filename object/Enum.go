package object

type NotifyType int

const (
	DingTalkRobot NotifyType = iota
)

type ActionType int

const (
	WindowsService ActionType = iota
	IISAppPool
)

type OprType int

const (
	Restart OprType = iota
)

type TaskType int

const (
	Int TaskType = iota
	CrmDzXfTest
	Health
	WebState
)

var TaskTypeList []TaskType

func init() {
	TaskTypeList = make([]TaskType, 0)
	TaskTypeList = append(TaskTypeList, Int)
	TaskTypeList = append(TaskTypeList, CrmDzXfTest)
	TaskTypeList = append(TaskTypeList, Health)
	TaskTypeList = append(TaskTypeList, WebState)
}
