package protos

type CafeTravelResponse struct {
	ResponsePacket
	Protocol Protocol
	FriendDB FriendDB
	CafeDBs  []CafeDB
}
