package protos

type EquipmentBatchGrowthRequest struct {
	RequestPacket
	Protocol                       Protocol
	EquipmentBatchGrowthRequestDBs []EquipmentBatchGrowthRequestDB
	GearTierUpRequestDB            GearTierUpRequestDB
}
