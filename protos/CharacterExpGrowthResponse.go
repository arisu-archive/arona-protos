package protos

type CharacterExpGrowthResponse struct {
	ResponsePacket
	Protocol          Protocol
	CharacterDB       CharacterDB
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
}
