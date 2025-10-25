package protos

type MissionRewardRequest struct {
	RequestPacket
	Protocol         Protocol
	MissionUniqueId  int64
	ProgressServerId int64
	EventContentId   *int64
}
