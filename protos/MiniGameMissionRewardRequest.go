package protos

type MiniGameMissionRewardRequest struct {
	RequestPacket
	Protocol         Protocol
	MissionUniqueId  int64
	ProgressServerId int64
	EventContentId   int64
}
