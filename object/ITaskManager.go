package object

type ITaskManager interface {
	GetIdList() []string
	GetTask(id string) ITask
	GetChRegister() chan<- ITask
	GetChUnregister() chan<- string
}
