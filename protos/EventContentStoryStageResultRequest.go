package protos

type EventContentStoryStageResultRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	StageUniqueId  int64
}
