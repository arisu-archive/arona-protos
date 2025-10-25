package protos

type MiniGameDefenseEnterBattleRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	StageId        int64
}
