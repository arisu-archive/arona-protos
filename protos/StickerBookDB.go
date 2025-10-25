package protos

type StickerBookDB struct {
	AccountId        int64
	UnusedStickerDBs []StickerDB
	UsedStickerDBs   []StickerDB
}
