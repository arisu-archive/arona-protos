package protos

type EventContentCollectionListResponse struct {
	ResponsePacket
	Protocol                Protocol
	EventContentUnlockCGDBs []EventContentCollectionDB
}
