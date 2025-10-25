package protos

type FriendBlockRequest struct {
	RequestPacket
	Protocol        Protocol
	TargetAccountId int64
}
