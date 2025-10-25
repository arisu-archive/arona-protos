package protos

type CampaignPurchasePlayCountHardStageResponse struct {
	ResponsePacket
	Protocol               Protocol
	AccountCurrencyDB      AccountCurrencyDB
	CampaignStageHistoryDB CampaignStageHistoryDB
}
