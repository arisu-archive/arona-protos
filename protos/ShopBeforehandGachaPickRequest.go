package protos

type ShopBeforehandGachaPickRequest struct {
	RequestPacket
	Protocol     Protocol
	ShopUniqueId int64
	GoodsId      int64
	TargetIndex  int64
}
