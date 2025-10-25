package protos

type MultiFloorRaidSyncResponse struct {
	ResponsePacket
	Protocol          Protocol
	MultiFloorRaidDBs []MultiFloorRaidDB
}
