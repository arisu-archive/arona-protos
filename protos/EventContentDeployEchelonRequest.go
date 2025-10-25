package protos

type EventContentDeployEchelonRequest struct {
	RequestPacket
	Protocol         Protocol
	EventContentId   int64
	StageUniqueId    int64
	DeployedEchelons []HexaUnit
}
