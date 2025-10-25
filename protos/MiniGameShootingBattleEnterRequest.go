package protos

type MiniGameShootingBattleEnterRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	UniqueId       int64
}
