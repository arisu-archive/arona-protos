package protos

type CafeOpenRequest struct {
	RequestPacket
	Protocol Protocol
	CafeId   int64
}
