package protos

type CafeRankUpResponse struct {
	ResponsePacket
	Protocol        Protocol
	CafeDB          CafeDB
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
}
