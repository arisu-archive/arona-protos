package protos

type EventContentShopBuyRefreshMerchandiseResponse struct {
	ResponsePacket
	Protocol                  Protocol
	AccountCurrencyDB         AccountCurrencyDB
	ConsumeResultDB           ConsumeResultDB
	ParcelResultDB            ParcelResultDB
	MailDB                    MailDB
	ShopProductDB             []ShopProductDB
	EventContentCollectionDBs []EventContentCollectionDB
}
