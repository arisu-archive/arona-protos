package protos

type ManagementProtocolLockListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ProtocolLockDBs []ProtocolLockDB `json:",omitempty,omitzero"`
}
