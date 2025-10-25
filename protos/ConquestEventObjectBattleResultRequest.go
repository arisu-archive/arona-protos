package protos

type ConquestEventObjectBattleResultRequest struct {
	RequestPacket
	Protocol           Protocol
	EventContentId     int64
	ConquestObjectDBId int64
	BattleSummary      BattleSummary
}
