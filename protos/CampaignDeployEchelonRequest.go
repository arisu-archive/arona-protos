package protos

type CampaignDeployEchelonRequest struct {
	RequestPacket
	Protocol         Protocol
	StageUniqueId    int64
	DeployedEchelons []HexaUnit
}
