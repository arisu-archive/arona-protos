package protos

type WorldRaidEnterBattleResponse struct {
	ResponsePacket
	Protocol           Protocol
	RaidBattleDB       RaidBattleDB
	AssistCharacterDBs []AssistCharacterDB
}
