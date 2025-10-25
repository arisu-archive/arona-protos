package protos

type RaidSeasonRewardResponse struct {
	ResponsePacket
	Protocol         Protocol
	ParcelResultDB   ParcelResultDB
	ReceiveRewardIds []int64
}
