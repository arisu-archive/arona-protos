package protos

type RaidEnterBattleResponse struct {
	ResponsePacket
	Protocol          Protocol
	RaidDB            RaidDB
	RaidBattleDB      RaidBattleDB
	AccountCurrencyDB AccountCurrencyDB
	AssistCharacterDB AssistCharacterDB
}
