package protos

type BattlePassBuyLevelRequest struct {
	RequestPacket
	Protocol                Protocol
	BattlePassId            int64
	BattlePassBuyLevelCount int32
}
