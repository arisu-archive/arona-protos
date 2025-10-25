package protos

type EventContentReceiveStageTotalRewardResponse struct {
	ResponsePacket
	Protocol               Protocol
	EventContentId         int64
	AlreadyReceiveRewardId []int64
	ParcelResultDB         ParcelResultDB
}
