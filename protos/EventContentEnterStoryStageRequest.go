package protos

type EventContentEnterStoryStageRequest struct {
	RequestPacket
	Protocol       Protocol
	StageUniqueId  int64
	EventContentId int64
}
