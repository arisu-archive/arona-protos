package protos

type EliminateRaidGiveUpResponse struct {
	ResponsePacket
	Protocol       Protocol
	Tier           int32
	RaidGiveUpDB   RaidGiveUpDB
	ParcelResultDB ParcelResultDB
}
