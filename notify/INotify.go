package notify

type iNotify interface {
	GetId() string
	SendMsg(msg string) error
}
