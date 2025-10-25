package protos

type CampaignEndTurnResponse struct {
	ResponsePacket
	Protocol          Protocol
	SaveDataDB        CampaignMainStageSaveDB
	AccountCurrencyDB AccountCurrencyDB
}
