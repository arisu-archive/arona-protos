package protos

type BattlePassGetInfoRequest struct {
	RequestPacket
	Protocol     Protocol
	BattlePassId int64
}
