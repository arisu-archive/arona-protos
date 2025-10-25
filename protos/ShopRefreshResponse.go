package protos

type ShopRefreshResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	ShopInfoDB     ShopInfoDB
}
