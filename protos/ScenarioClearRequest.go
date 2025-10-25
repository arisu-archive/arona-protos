package protos

type ScenarioClearRequest struct {
	RequestPacket
	Protocol      Protocol
	ScenarioId    int64
	BattleSummary BattleSummary
}
