package protos

type BillingPurchaseCashShopVerifyByNexonRequest struct {
	RequestPacket
	NpSN           int64 `json:",omitempty,omitzero"`
	StampToken     string
	ShopCashId     int64 `json:",omitempty,omitzero"`
	VirtualPayment bool  `json:",omitempty,omitzero"`
	CurrencyCode   string
	CurrencyValue  int64 `json:",omitempty,omitzero"`
}
