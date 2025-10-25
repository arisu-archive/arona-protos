package protos

type CharacterSetCostumeRequest struct {
	RequestPacket
	Protocol          Protocol
	CharacterUniqueId int64
	CostumeIdToSet    *int64
}
