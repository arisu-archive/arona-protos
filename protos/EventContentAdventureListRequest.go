package protos

type EventContentAdventureListRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
