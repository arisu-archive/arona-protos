package protos

type ClanMemberRequest struct {
	RequestPacket
	Protocol        Protocol
	ClanDBId        int64
	MemberAccountId int64
}
