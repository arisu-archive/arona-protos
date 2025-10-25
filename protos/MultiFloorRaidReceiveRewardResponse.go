package protos

type MultiFloorRaidReceiveRewardResponse struct {
	ResponsePacket
	Protocol         Protocol
	MultiFloorRaidDB MultiFloorRaidDB
	ParcelResultDB   ParcelResultDB
}
