package protos

type ShopBuyGachaRequest struct {
	RequestPacket
	Protocol     Protocol
	GoodsId      int64
	ShopUniqueId int64
}
