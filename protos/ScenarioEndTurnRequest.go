package protos

type ScenarioEndTurnRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
