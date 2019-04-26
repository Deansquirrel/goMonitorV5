package action

type OprType int

const (
	Restart OprType = iota
)

type IAction interface {
	//操作接口
	Do(oprType OprType, id string) error
}
