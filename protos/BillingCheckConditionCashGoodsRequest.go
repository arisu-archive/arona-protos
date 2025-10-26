package protos

type BillingCheckConditionCashGoodsRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	User_id string `json:",omitempty,omitzero"`
	Product_id int64 `json:",omitempty,omitzero"`
}
