package protos

type CharacterSetCostumeResponse struct {
	ResponsePacket
	Protocol       Protocol
	SetCostumeDB   CostumeDB
	UnsetCostumeDB CostumeDB
}
