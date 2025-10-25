package protos

type EquipmentItemEquipRequest struct {
	RequestPacket
	Protocol           Protocol
	CharacterServerId  int64
	EquipmentServerIds []int64
	EquipmentServerId  int64
	SlotIndex          int32
}
