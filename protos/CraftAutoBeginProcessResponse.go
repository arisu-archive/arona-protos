package protos

type CraftAutoBeginProcessResponse struct {
	ResponsePacket
	Protocol       Protocol
	CraftInfoDBs   []CraftInfoDB
	ParcelResultDB ParcelResultDB
}
