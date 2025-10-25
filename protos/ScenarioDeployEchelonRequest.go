package protos

type ScenarioDeployEchelonRequest struct {
	RequestPacket
	Protocol         Protocol
	StageUniqueId    int64
	DeployedEchelons []HexaUnit
}
