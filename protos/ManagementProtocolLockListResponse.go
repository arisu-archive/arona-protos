package protos

type ManagementProtocolLockListResponse struct {
	ResponsePacket
	Protocol        Protocol
	ProtocolLockDBs []ProtocolLockDB
}
