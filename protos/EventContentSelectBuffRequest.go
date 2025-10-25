package protos

type EventContentSelectBuffRequest struct {
	RequestPacket
	Protocol       Protocol
	SelectedBuffId int64
}
