package protos

type BillingPurchaseListByNexonResponse struct {
	ResponsePacket
	Protocol                   Protocol
	CountList                  []PurchaseCountDB
	OrderList                  []PurchaseOrderDB
	MonthlyProductList         []MonthlyProductPurchaseDB
	ProductMonthlyIdInMailList []int64
	GachaTicketItemIdList      []int64
	BlockedProductDBs          []BlockedProductDB
	BattlePassProductList      []BattlePassProductPurchaseDB
	BattlePassIdInMailList     []int64
	IsTeenage                  bool
}
