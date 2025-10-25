package protos

type FriendAcceptFriendRequestResponse struct {
	ResponsePacket
	Protocol                 Protocol
	FriendDBs                []FriendDB
	SentRequestFriendDBs     []FriendDB
	ReceivedRequestFriendDBs []FriendDB
	BlockedUserDBs           []FriendDB
}
