package protos

type ShopBeforehandGachaRunRequest struct {
	RequestPacket
	Protocol     Protocol
	ShopUniqueId int64
	GoodsId      int64
}
