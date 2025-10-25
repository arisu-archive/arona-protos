package protos

type ScenarioConfirmMainStageRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
