package protos

type BillingPurchaseCashShopVerifyByNexonRequest struct {
	RequestPacket
	NpSN           int64  `json:",omitempty,omitzero"`
	StampToken     string `json:",omitempty,omitzero"`
	ShopCashId     int64  `json:",omitempty,omitzero"`
	VirtualPayment bool   `json:",omitempty,omitzero"`
	CurrencyCode   string `json:",omitempty,omitzero"`
	CurrencyValue  int64  `json:",omitempty,omitzero"`
}
