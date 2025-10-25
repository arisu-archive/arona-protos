package protos

type CafeAckResponse struct {
	ResponsePacket
	Protocol Protocol
	CafeDB   CafeDB
}
