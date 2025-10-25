package protos

type NotificationLobbyCheckResponse struct {
	ResponsePacket
	Protocol               Protocol
	UnreadMailCount        int64
	EventRewardIncreaseDBs []EventRewardIncreaseDB
}
