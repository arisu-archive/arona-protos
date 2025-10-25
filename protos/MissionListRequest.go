package protos

type MissionListRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId *int64
}
