package protos

type ConquestReceiveRewardsResponse struct {
	ResponsePacket
	Protocol        Protocol
	ParcelResultDB  ParcelResultDB
	ConquestInfoDB  ConquestInfoDB
	ConquestTileDBs []ConquestTileDB
}
