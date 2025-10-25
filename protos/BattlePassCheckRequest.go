package protos

type BattlePassCheckRequest struct {
	RequestPacket
	Protocol     Protocol
	BattlePassId int64
}
