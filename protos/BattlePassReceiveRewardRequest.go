package protos

type BattlePassReceiveRewardRequest struct {
	RequestPacket
	Protocol     Protocol
	BattlePassId int64
}
