package protos

type EventContentPurchasePlayCountHardStageResponse struct {
	ResponsePacket
	Protocol               Protocol
	AccountCurrencyDB      AccountCurrencyDB
	CampaignStageHistoryDB CampaignStageHistoryDB
}
