package protos

type RaidCreateBattleResponse struct {
	ResponsePacket
	Protocol          Protocol
	RaidDB            RaidDB
	RaidBattleDB      RaidBattleDB
	AccountCurrencyDB AccountCurrencyDB
	AssistCharacterDB AssistCharacterDB
}
