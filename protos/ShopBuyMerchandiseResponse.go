package protos

type ShopBuyMerchandiseResponse struct {
	ResponsePacket
	Protocol          Protocol
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
	ParcelResultDB    ParcelResultDB
	MailDB            MailDB
	ShopProductDB     ShopProductDB
}
