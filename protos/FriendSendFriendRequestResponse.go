package protos

type FriendSendFriendRequestResponse struct {
	ResponsePacket
	Protocol                 Protocol
	FriendDBs                []FriendDB
	SentRequestFriendDBs     []FriendDB
	ReceivedRequestFriendDBs []FriendDB
	BlockedUserDBs           []FriendDB
}
