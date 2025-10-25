package protos

type EventContentLocationGetInfoRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
