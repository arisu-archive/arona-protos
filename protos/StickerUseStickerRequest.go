package protos

type StickerUseStickerRequest struct {
	RequestPacket
	Protocol        Protocol
	StickerUniqueId int64
}
