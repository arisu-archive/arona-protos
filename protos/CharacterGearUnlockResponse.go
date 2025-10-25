package protos

type CharacterGearUnlockResponse struct {
	ResponsePacket
	Protocol    Protocol
	GearDB      GearDB
	CharacterDB CharacterDB
}
