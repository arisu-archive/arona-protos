package protos

type FriendRemoveRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
