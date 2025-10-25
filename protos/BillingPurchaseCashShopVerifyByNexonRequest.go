package protos

type BillingPurchaseCashShopVerifyByNexonRequest struct {
	RequestPacket
	Protocol       Protocol
	NpSN           int64
	StampToken     string
	ShopCashId     int64
	VirtualPayment bool
	CurrencyCode   string
	CurrencyValue  int64
}
