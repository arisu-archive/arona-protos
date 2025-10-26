package protos

type BillingPurchaseListByNexonRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	IsTeenage bool `json:",omitempty,omitzero"`
}
