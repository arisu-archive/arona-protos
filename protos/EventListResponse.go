package protos

type EventListResponse struct {
	ResponsePacket
	Protocol     Protocol
	EventInfoDBs []EventInfoDB
}
