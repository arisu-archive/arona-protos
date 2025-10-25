package protos

type ScenarioGroupHistoryUpdateRequest struct {
	RequestPacket
	Protocol              Protocol
	ScenarioGroupUniqueId int64
	ScenarioType          int64
}
