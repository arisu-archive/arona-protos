package protos

type EventContentDeployEchelonResponse struct {
	ResponsePacket
	Protocol   Protocol
	SaveDataDB EventContentMainStageSaveDB
}
