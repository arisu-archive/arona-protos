package protos

type FriendUnblockResponse struct {
	ResponsePacket
	Protocol                 Protocol
	FriendDBs                []FriendDB
	SentRequestFriendDBs     []FriendDB
	ReceivedRequestFriendDBs []FriendDB
	BlockedUserDBs           []FriendDB
}
