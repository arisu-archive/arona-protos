package protos

type BattlePassMissionListRequest struct {
	RequestPacket
	Protocol     Protocol
	BattlePassId int64
}
