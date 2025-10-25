package protos

type ClanAutoJoinResponse struct {
	ResponsePacket
	Protocol     Protocol
	IrcConfig    IrcServerConfig
	ClanDB       ClanDB
	ClanMemberDB ClanMemberDB
}
