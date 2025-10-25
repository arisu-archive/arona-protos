package protos

type FriendSearchResponse struct {
	ResponsePacket
	Protocol     Protocol
	SearchResult []FriendDB
}
