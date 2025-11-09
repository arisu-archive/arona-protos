package protos

type BillingPurchaseListByNexonResponse struct {
	ResponsePacket
	CountList                  []PurchaseCountDB             `json:",omitempty,omitzero"`
	OrderList                  []PurchaseOrderDB             `json:",omitempty,omitzero"`
	MonthlyProductList         []MonthlyProductPurchaseDB    `json:",omitempty,omitzero"`
	ProductMonthlyIdInMailList []int64                       `json:",omitempty,omitzero"`
	GachaTicketItemIdList      []int64                       `json:",omitempty,omitzero"`
	BlockedProductDBs          []BlockedProductDB            `json:",omitempty,omitzero"`
	BattlePassProductList      []BattlePassProductPurchaseDB `json:",omitempty,omitzero"`
	BattlePassIdInMailList     []int64                       `json:",omitempty,omitzero"`
	IsTeenage                  bool                          `json:",omitempty,omitzero"`
}
