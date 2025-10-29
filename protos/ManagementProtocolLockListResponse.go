package protos

type ManagementProtocolLockListResponse struct {
	ResponsePacket
	ProtocolLockDBs []ProtocolLockDB `json:",omitempty,omitzero"`
}
