package protos

type FriendUnblockRequest struct {
	RequestPacket
	Protocol        Protocol
	TargetAccountId int64
}
