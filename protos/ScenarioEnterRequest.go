package protos

type ScenarioEnterRequest struct {
	RequestPacket
	Protocol   Protocol
	ScenarioId int64
}
