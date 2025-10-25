package protos

type ShopBuyRefreshMerchandiseResponse struct {
	ResponsePacket
	Protocol          Protocol
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
	ParcelResultDB    ParcelResultDB
	ShopProductDB     []ShopProductDB
	MailDB            MailDB
}
