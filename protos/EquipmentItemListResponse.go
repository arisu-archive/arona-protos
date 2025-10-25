package protos

type EquipmentItemListResponse struct {
	ResponsePacket
	Protocol     Protocol
	EquipmentDBs []EquipmentDB
}
