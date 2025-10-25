package protos

type ClanCreateResponse struct {
	ResponsePacket
	Protocol          Protocol
	ClanDB            ClanDB
	ClanMemberDB      ClanMemberDB
	AccountCurrencyDB AccountCurrencyDB
}
