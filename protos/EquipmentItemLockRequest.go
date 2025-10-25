package protos

type EquipmentItemLockRequest struct {
	RequestPacket
	Protocol       Protocol
	TargetServerId int64
	IsLocked       bool
}
