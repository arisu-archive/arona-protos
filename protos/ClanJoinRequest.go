package protos

type ClanJoinRequest struct {
	RequestPacket
	Protocol Protocol
	ClanDBId int64
}
