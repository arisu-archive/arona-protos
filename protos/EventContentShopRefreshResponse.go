package protos

type EventContentShopRefreshResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	ShopInfoDB     ShopInfoDB
}
