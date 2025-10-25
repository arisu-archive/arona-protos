package protos

type FriendGetFriendDetailedInfoRequest struct {
	RequestPacket
	Protocol        Protocol
	FriendAccountId int64
}
