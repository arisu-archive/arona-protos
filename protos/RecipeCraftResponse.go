package protos

type RecipeCraftResponse struct {
	ResponsePacket
	Protocol                 Protocol
	ParcelResultDB           ParcelResultDB
	EquipmentConsumeResultDB ConsumeResultDB
	ItemConsumeResultDB      ConsumeResultDB
}
