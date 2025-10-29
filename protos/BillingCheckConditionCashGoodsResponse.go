package protos

type BillingCheckConditionCashGoodsResponse struct {
	ResponsePacket
	Result bool `json:",omitempty,omitzero"`
}
