package protos

type FriendListByIdsResponse struct {
	ResponsePacket
	Protocol   Protocol
	ListResult []FriendDB
}
