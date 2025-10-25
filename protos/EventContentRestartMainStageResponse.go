package protos

type EventContentRestartMainStageResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	SaveDataDB     EventContentMainStageSaveDB
}
