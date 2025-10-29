package protos

type BillingPurchaseListByNexonRequest struct {
	RequestPacket
	IsTeenage bool `json:",omitempty,omitzero"`
}
