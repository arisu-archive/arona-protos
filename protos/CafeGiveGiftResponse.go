package protos

type CafeGiveGiftResponse struct {
	ResponsePacket
	Protocol        Protocol
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
}
