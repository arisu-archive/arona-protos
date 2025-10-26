package protos

type BillingCheckConditionCashGoodsResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Result bool `json:",omitempty,omitzero"`
}
