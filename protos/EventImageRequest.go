package protos

type EventImageRequest struct {
	RequestPacket
	Protocol Protocol
	EventId  int64
}
