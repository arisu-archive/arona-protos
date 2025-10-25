package protos

type ClanKickRequest struct {
	RequestPacket
	Protocol        Protocol
	MemberAccountId int64
}
