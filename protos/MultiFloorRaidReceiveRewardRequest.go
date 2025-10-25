package protos

type MultiFloorRaidReceiveRewardRequest struct {
	RequestPacket
	Protocol         Protocol
	SeasonId         int64
	RewardDifficulty int32
}
