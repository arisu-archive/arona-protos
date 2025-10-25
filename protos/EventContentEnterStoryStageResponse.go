package protos

type EventContentEnterStoryStageResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	SaveDataDB     EventContentStoryStageSaveDB
}
