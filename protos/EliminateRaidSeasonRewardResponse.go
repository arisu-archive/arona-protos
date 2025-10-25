package protos

type EliminateRaidSeasonRewardResponse struct {
	ResponsePacket
	Protocol         Protocol
	ParcelResultDB   ParcelResultDB
	ReceiveRewardIds []int64
}
