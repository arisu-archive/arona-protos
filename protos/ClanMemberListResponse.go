package protos

type ClanMemberListResponse struct {
	ResponsePacket
	Protocol      Protocol
	ClanDB        ClanDB
	ClanMemberDBs []ClanMemberDB
}
