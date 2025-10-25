package protos

type StickerUseStickerResponse struct {
	ResponsePacket
	Protocol       Protocol
	StickerBookDB  StickerBookDB
	ParcelResultDB ParcelResultDB
}
