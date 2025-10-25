package protos

type CafeAckRequest struct {
	RequestPacket
	Protocol Protocol
	CafeDBId int64
}
