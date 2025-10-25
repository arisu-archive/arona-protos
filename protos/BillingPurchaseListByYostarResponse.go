package protos

type BillingPurchaseListByYostarResponse struct {
	ResponsePacket
	Protocol              Protocol
	CountList             []PurchaseCountDB
	OrderList             []PurchaseOrderDB
	MonthlyProductList    []MonthlyProductPurchaseDB
	BlockedProductDBs     []BlockedProductDB
	BattlePassProductList []BattlePassProductPurchaseDB
}
