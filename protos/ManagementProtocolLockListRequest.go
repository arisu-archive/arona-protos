package protos

type ManagementProtocolLockListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
