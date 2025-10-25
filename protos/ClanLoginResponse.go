package protos

type ClanLoginResponse struct {
	ResponsePacket
	Protocol            Protocol
	AccountClanDB       ClanDB
	AccountClanMemberDB ClanMemberDB
	ClanAssistSlotDBs   []ClanAssistSlotDB
}
