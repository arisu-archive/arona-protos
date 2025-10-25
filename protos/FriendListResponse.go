package protos

type FriendListResponse struct {
	ResponsePacket
	Protocol                 Protocol
	IdCardBackgroundDBs      []IdCardBackgroundDB
	FriendDBs                []FriendDB
	SentRequestFriendDBs     []FriendDB
	ReceivedRequestFriendDBs []FriendDB
	BlockedUserDBs           []FriendDB
}
