package protos

type EliminateRaidLimitedRewardResponse struct {
	ResponsePacket
	Protocol         Protocol
	ParcelResultDB   ParcelResultDB
	ReceiveRewardIds []int64
}
