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
	IntD
	CrmDzXfTest
	Health
	WebState
)
