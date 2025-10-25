package protos

type EventContentConfirmMainStageResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	SaveDataDB     EventContentMainStageSaveDB
}
