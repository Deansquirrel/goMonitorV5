package notify

type iNotify interface {
	SendMsg(msg string) error
}
