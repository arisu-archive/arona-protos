package protos

type EventRewardIncreaseResponse struct {
	ResponsePacket
	Protocol               Protocol
	EventRewardIncreaseDBs []EventRewardIncreaseDB
}
