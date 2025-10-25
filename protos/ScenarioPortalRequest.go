package protos

type ScenarioPortalRequest struct {
	RequestPacket
	Protocol        Protocol
	StageUniqueId   int64
	EchelonEntityId int64
}
