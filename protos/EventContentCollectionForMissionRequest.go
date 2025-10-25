package protos

type EventContentCollectionForMissionRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
