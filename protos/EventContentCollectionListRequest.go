package protos

type EventContentCollectionListRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	GroupId        *int64
}
