package protos

type BillingPurchaseCashShopVerifyByNexonResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResult ParcelResultDB `json:",omitempty,omitzero"`
	MailDB MailDB `json:",omitempty,omitzero"`
	CountList []PurchaseCountDB `json:",omitempty,omitzero"`
	PurchaseCount int32 `json:",omitempty,omitzero"`
	MonthlyProductList []MonthlyProductPurchaseDB `json:",omitempty,omitzero"`
	ProductMonthlyIdInMailList []int64 `json:",omitempty,omitzero"`
	GachaTicketItemIdList []int64 `json:",omitempty,omitzero"`
	ShopId string `json:",omitempty,omitzero"`
	ItemPrice float64 `json:",omitempty,omitzero"`
	Currency string `json:",omitempty,omitzero"`
	StampId string `json:",omitempty,omitzero"`
	BattlePassIdInMailList []int64 `json:",omitempty,omitzero"`
}
