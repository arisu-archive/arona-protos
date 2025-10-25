package protos

type ClanMemberListRequest struct {
	RequestPacket
	Protocol Protocol
	ClanDBId int64
}
