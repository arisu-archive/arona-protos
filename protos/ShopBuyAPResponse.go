package protos

type ShopBuyAPResponse struct {
	ResponsePacket
	Protocol          Protocol
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
	ParcelResultDB    ParcelResultDB
	MailDB            MailDB
	ShopProductDB     ShopProductDB
}
