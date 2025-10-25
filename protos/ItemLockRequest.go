package protos

type ItemLockRequest struct {
	RequestPacket
	Protocol       Protocol
	TargetServerId int64
	IsLocked       bool
}
