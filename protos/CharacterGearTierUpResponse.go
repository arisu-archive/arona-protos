package protos

type CharacterGearTierUpResponse struct {
	ResponsePacket
	Protocol        Protocol
	GearDB          GearDB
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
}
