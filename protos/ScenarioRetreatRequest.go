package protos

type ScenarioRetreatRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
}
