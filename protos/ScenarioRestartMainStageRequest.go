package protos

type ScenarioRestartMainStageRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
