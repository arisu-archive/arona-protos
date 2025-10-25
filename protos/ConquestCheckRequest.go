package protos

type ConquestCheckRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
