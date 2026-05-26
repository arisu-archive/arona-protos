package protos

type AccountCallNameRequest struct {
	RequestPacket
	CallName         string
	CallNameKatakana string
	CallNameKorean   string
}
