package protos

type EventContentEnterSubStageRequest struct {
	RequestPacket
	Protocol                    Protocol
	EventContentId              int64
	StageUniqueId               int64
	LastEnterStageEchelonNumber int64
}
