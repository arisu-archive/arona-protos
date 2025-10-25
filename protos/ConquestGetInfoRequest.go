package protos

type ConquestGetInfoRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
