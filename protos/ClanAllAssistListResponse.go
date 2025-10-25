package protos

type ClanAllAssistListResponse struct {
	ResponsePacket
	Protocol                      Protocol
	AssistCharacterDBs            []AssistCharacterDB
	AssistCharacterRentHistoryDBs []ClanAssistRentHistoryDB
	ClanDBId                      int64
}
