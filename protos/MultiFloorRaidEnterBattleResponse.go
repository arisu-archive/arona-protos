package protos

type MultiFloorRaidEnterBattleResponse struct {
	ResponsePacket
	Protocol           Protocol
	AssistCharacterDBs []AssistCharacterDB
}
