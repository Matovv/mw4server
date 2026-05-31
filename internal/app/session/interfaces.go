package session

type ClientSender interface {
	GetId() uint64
    Send(any) error
}