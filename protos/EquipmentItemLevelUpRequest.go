package protos

type EquipmentItemLevelUpRequest struct {
	RequestPacket
	Protocol         Protocol
	TargetServerId   int64
	ConsumeServerIds []int64
	ConsumeRequestDB ConsumeRequestDB
}
