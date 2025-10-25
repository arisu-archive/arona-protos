package protos

type BillingPurchaseCashShopVerifyByNexonResponse struct {
	ResponsePacket
	Protocol                   Protocol
	ParcelResult               ParcelResultDB
	MailDB                     MailDB
	CountList                  []PurchaseCountDB
	PurchaseCount              int32
	MonthlyProductList         []MonthlyProductPurchaseDB
	ProductMonthlyIdInMailList []int64
	GachaTicketItemIdList      []int64
	ShopId                     string
	ItemPrice                  float64
	Currency                   string
	StampId                    string
	BattlePassIdInMailList     []int64
}
