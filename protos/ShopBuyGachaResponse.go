package protos

type ShopBuyGachaResponse struct {
	ResponsePacket
	Protocol        Protocol
	ConsumeResultDB ConsumeResultDB
	ParcelResultDB  ParcelResultDB
}
