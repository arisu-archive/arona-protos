package protos

type ConquestEventObjectBattleResultResponse struct {
	ResponsePacket
	Protocol                     Protocol
	ParcelResultDB               ParcelResultDB
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	ConquestInfoDB               ConquestInfoDB
	ConquestTileDB               ConquestTileDB
	DisplayInfos                 []ConquestDisplayInfo
}
