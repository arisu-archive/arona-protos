package protos

type EquipmentItemSellRequest struct {
	RequestPacket
	Protocol        Protocol
	TargetServerIds []int64
}
