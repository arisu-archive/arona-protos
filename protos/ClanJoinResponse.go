package protos

type ClanJoinResponse struct {
	ResponsePacket
	Protocol     Protocol
	IrcConfig    IrcServerConfig
	ClanDB       ClanDB
	ClanMemberDB ClanMemberDB
}
