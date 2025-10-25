package protos

type EventContentScenarioGroupHistoryUpdateRequest struct {
	RequestPacket
	Protocol              Protocol
	ScenarioGroupUniqueId int64
	ScenarioType          int64
	EventContentId        int64
}
