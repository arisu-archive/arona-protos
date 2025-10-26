package protos

type FriendCancelFriendRequestRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
