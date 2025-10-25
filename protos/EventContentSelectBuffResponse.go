package protos

type EventContentSelectBuffResponse struct {
	ResponsePacket
	Protocol   Protocol
	SaveDataDB EventContentMainStageSaveDB
}
