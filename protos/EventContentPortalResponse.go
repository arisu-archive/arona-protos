package protos

type EventContentPortalResponse struct {
	ResponsePacket
	Protocol   Protocol
	SaveDataDB EventContentMainStageSaveDB
}
