package protos

type MultiFloorRaidSyncRequest struct {
	RequestPacket
	Protocol Protocol
	SeasonId *int64
}
