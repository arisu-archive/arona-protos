package protos

type CampaignHealResponse struct {
	ResponsePacket
	Protocol          Protocol
	AccountCurrencyDB AccountCurrencyDB
	SaveDataDB        CampaignMainStageSaveDB
}
