package protos

type CharacterGearListResponse struct {
	ResponsePacket
	Protocol Protocol
	GearDBs  []GearDB
}
