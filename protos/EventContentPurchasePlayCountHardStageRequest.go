package protos

type EventContentPurchasePlayCountHardStageRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	StageUniqueId  int64
}
