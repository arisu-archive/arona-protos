package protos

type ContentSweepMultiSweepResponse struct {
	ResponsePacket
	Protocol                 Protocol
	ClearParcels             [][]ParcelInfo
	BonusParcels             []ParcelInfo
	EventContentBonusParcels [][]ParcelInfo
	ParcelResult             ParcelResultDB
	CampaignStageHistoryDBs  []CampaignStageHistoryDB
}
