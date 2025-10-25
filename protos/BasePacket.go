package protos

type BasePacket struct {
	SessionKey SessionKey
	Protocol   Protocol
	AccountId  int64
}
