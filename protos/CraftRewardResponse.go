package protos

type CraftRewardResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	CraftInfos     []CraftInfoDB
}
