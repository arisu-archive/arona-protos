package protos

type ScenarioEnterMainStageRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
