package protos

type AccountNicknameRequest struct {
	RequestPacket
	Protocol Protocol
	Nickname string
}
