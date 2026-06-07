package session

type ClientSender interface {
	GetId() uint64
    SendPacket(
		packetType string,
		payload any,
	) error
}