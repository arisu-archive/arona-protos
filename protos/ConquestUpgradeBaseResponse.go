package protos

type ConquestUpgradeBaseResponse struct {
	ResponsePacket
	Protocol                     Protocol
	UpgradeRewards               []ParcelInfo
	ParcelResultDB               ParcelResultDB
	ConquestTileDB               ConquestTileDB
	ConquestInfoDB               ConquestInfoDB
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	DisplayInfos                 []ConquestDisplayInfo
}
