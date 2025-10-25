package protos

type ConquestManageBaseResponse struct {
	ResponsePacket
	Protocol                     Protocol
	ClearParcels                 [][]ParcelInfo
	ConquerBonusParcels          [][]ParcelInfo
	BonusParcels                 []ParcelInfo
	ParcelResultDB               ParcelResultDB
	ConquestInfoDB               ConquestInfoDB
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	DisplayInfos                 []ConquestDisplayInfo
}
