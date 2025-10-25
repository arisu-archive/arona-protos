package protos

type ClanPermitResponse struct {
	ResponsePacket
	Protocol     Protocol
	ClanDB       ClanDB
	ClanMemberDB ClanMemberDB
}
