package protos

type CharacterSetFavoritesRequest struct {
	RequestPacket
	Protocol            Protocol
	ActivateByServerIds map[int64]bool
}
