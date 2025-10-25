package protos

type EventContentReceiveStageTotalRewardRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
}
