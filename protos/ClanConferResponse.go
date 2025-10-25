package protos

type ClanConferResponse struct {
	ResponsePacket
	Protocol                Protocol
	ClanMemberDB            ClanMemberDB
	AccountClanMemberDB     ClanMemberDB
	ClanDB                  ClanDB
	ClanMemberDescriptionDB ClanMemberDescriptionDB
}
