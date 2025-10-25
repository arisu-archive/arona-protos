package protos

type ConquestErosionBattleStartRequest struct {
	RequestPacket
	Protocol           Protocol
	EventContentId     int64
	ConquestObjectDBId int64
	UseManageEchelon   bool
	EchelonNumber      int64
	ClanAssistUseInfo  ClanAssistUseInfo
}
