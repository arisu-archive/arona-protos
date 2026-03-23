package protos

type BillingPurchaseCashShopVerifyByNexonResponse struct {
	ResponsePacket
	ParcelResult               *ParcelResultDB `json:",omitempty,omitzero"`
	MailDB                     *MailDB         `json:",omitempty,omitzero"`
	CountList                  []PurchaseCountDB
	PurchaseCount              int32 `json:",omitempty,omitzero"`
	MonthlyProductList         []MonthlyProductPurchaseDB
	ProductMonthlyIdInMailList []int64
	GachaTicketItemIdList      []int64
	ShopId                     string
	ItemPrice                  float64 `json:",omitempty,omitzero"`
	Currency                   string
	StampId                    string
	BattlePassIdInMailList     []int64
}
