package protos

type EquipmentBatchGrowthResponse struct {
	ResponsePacket
	Protocol        Protocol
	EquipmentDBs    []EquipmentDB
	GearDB          GearDB
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
}
