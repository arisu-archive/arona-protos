package protos

type FriendBlockResponse struct {
	ResponsePacket
	Protocol                 Protocol
	FriendDBs                []FriendDB
	SentRequestFriendDBs     []FriendDB
	ReceivedRequestFriendDBs []FriendDB
	BlockedUserDBs           []FriendDB
}
