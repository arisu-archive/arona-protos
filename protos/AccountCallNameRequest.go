package protos

type AccountCallNameRequest struct {
	RequestPacket
	Protocol         Protocol
	CallName         string
	CallNameKatakana string
	CallNameKorean   string
}
