package protos

type BattlePassBuyLevelResponse struct {
	ResponsePacket
	Protocol          Protocol
	BattlePassInfo    BattlePassInfoDB
	AccountCurrencyDB AccountCurrencyDB
}
