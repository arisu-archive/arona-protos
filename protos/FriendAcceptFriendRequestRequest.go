package protos

type FriendAcceptFriendRequestRequest struct {
	RequestPacket
	Protocol        Protocol
	TargetAccountId int64
}
