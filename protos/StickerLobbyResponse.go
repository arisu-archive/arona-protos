package protos

type StickerLobbyResponse struct {
	ResponsePacket
	Protocol           Protocol
	ReceivedStickerDBs []StickerDB
	StickerBookDB      StickerBookDB
}
