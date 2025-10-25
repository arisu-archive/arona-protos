package protos

type EventContentPermanentListResponse struct {
	ResponsePacket
	Protocol     Protocol
	PermanentDBs []EventContentPermanentDB
}
