package protos

type CharacterSetFavoritesResponse struct {
	ResponsePacket
	Protocol     Protocol
	CharacterDBs []CharacterDB
}
