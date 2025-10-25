package protos

type RaidGiveUpResponse struct {
	ResponsePacket
	Protocol       Protocol
	Tier           int32
	RaidGiveUpDB   RaidGiveUpDB
	ParcelResultDB ParcelResultDB
}
