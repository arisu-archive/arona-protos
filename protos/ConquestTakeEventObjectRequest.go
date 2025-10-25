package protos

type ConquestTakeEventObjectRequest struct {
	RequestPacket
	Protocol           Protocol
	EventContentId     int64
	ConquestObjectDBId int64
}
