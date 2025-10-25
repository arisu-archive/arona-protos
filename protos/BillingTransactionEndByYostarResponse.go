package protos

type BillingTransactionEndByYostarResponse struct {
	ResponsePacket
	Protocol              Protocol
	ParcelResult          ParcelResultDB
	MailDB                MailDB
	CountList             []PurchaseCountDB
	PurchaseCount         int32
	MonthlyProductList    []MonthlyProductPurchaseDB
	BattlePassInfo        BattlePassInfoDB
	BattlePassProductList []BattlePassProductPurchaseDB
}
