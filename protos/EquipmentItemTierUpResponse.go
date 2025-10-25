package protos

type EquipmentItemTierUpResponse struct {
	ResponsePacket
	Protocol        Protocol
	EquipmentDB     EquipmentDB
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
}
