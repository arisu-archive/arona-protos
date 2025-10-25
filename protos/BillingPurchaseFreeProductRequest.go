package protos

type BillingPurchaseFreeProductRequest struct {
	RequestPacket
	Protocol   Protocol
	ShopCashId int64
}
