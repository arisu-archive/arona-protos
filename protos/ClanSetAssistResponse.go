package protos

type ClanSetAssistResponse struct {
	ResponsePacket
	Protocol         Protocol
	ClanAssistSlotDB ClanAssistSlotDB
	ParcelResultDB   ParcelResultDB
	RewardInfo       ClanAssistRewardInfo
}
