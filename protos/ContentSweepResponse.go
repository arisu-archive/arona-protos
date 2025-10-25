package protos

type ContentSweepResponse struct {
	ResponsePacket
	Protocol                 Protocol
	ClearParcels             [][]ParcelInfo
	BonusParcels             []ParcelInfo
	EventContentBonusParcels [][]ParcelInfo
	ParcelResult             ParcelResultDB
	CampaignStageHistoryDB   CampaignStageHistoryDB
}
