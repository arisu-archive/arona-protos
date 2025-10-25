package protos

type CafeRankUpRequest struct {
	RequestPacket
	Protocol         Protocol
	AccountServerId  int64
	CafeDBId         int64
	ConsumeRequestDB ConsumeRequestDB
}
