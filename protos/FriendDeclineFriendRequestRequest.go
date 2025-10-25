package protos

type FriendDeclineFriendRequestRequest struct {
	RequestPacket
	Protocol        Protocol
	TargetAccountId int64
}
