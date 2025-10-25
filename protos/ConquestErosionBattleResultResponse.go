package protos

type ConquestErosionBattleResultResponse struct {
	ResponsePacket
	Protocol                     Protocol
	ParcelResultDB               ParcelResultDB
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	ConquestInfoDB               ConquestInfoDB
	DisplayInfos                 []ConquestDisplayInfo
}
