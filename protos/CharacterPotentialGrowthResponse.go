package protos

type CharacterPotentialGrowthResponse struct {
	ResponsePacket
	Protocol       Protocol
	CharacterDB    CharacterDB
	ParcelResultDB ParcelResultDB
}
