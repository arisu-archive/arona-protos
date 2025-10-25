package protos

type ScenarioSkipMainStageRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
