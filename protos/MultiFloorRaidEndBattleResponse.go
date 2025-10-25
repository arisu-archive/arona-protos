package protos

type MultiFloorRaidEndBattleResponse struct {
	ResponsePacket
	Protocol         Protocol
	MultiFloorRaidDB MultiFloorRaidDB
	ParcelResultDB   ParcelResultDB
}
