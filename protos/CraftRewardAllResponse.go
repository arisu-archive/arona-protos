package protos

type CraftRewardAllResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	CraftInfos     []CraftInfoDB
}
