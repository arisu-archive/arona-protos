package protos

type StickerLoginResponse struct {
	ResponsePacket
	Protocol      Protocol
	StickerBookDB StickerBookDB
}
