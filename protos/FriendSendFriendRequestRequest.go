package protos

type FriendSendFriendRequestRequest struct {
	RequestPacket
	Protocol        Protocol
	TargetAccountId int64
}
