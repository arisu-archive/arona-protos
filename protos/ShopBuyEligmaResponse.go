package protos

type ShopBuyEligmaResponse struct {
	ResponsePacket
	Protocol        Protocol
	ParcelResultDB  ParcelResultDB
	ConsumeResultDB ConsumeResultDB
	ShopProductDB   ShopProductDB
}
