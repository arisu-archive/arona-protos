package protos

type ConquestConquerResponse struct {
	ResponsePacket
	Protocol                     Protocol
	ParcelResultDB               ParcelResultDB
	ConquestTileDB               ConquestTileDB
	ConquestInfoDB               ConquestInfoDB
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	DisplayInfos                 []ConquestDisplayInfo
}
