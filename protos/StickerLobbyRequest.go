package protos

type StickerLobbyRequest struct {
	RequestPacket
	Protocol                Protocol
	AcquireStickerUniqueIds []int64
}
