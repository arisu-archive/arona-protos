package protos

type EventContentEnterMainStageResponse struct {
	ResponsePacket
	Protocol     Protocol
	SaveDataDB   EventContentMainStageSaveDB
	IsOnSubEvent bool
}
