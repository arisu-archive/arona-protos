package protos

type ConquestEventObjectBattleStartRequest struct {
	RequestPacket
	Protocol           Protocol
	EventContentId     int64
	ConquestObjectDBId int64
	EchelonNumber      int64
	ClanAssistUseInfo  ClanAssistUseInfo
}
