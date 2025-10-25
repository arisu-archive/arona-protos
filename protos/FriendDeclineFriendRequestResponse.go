package protos

type FriendDeclineFriendRequestResponse struct {
	ResponsePacket
	Protocol                 Protocol
	FriendDBs                []FriendDB
	SentRequestFriendDBs     []FriendDB
	ReceivedRequestFriendDBs []FriendDB
	BlockedUserDBs           []FriendDB
}
