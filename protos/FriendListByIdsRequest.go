package protos

type FriendListByIdsRequest struct {
	RequestPacket
	Protocol         Protocol
	TargetAccountIds []int64
}
