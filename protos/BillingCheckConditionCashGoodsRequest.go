package protos

type BillingCheckConditionCashGoodsRequest struct {
	RequestPacket
	User_id    string `json:",omitempty,omitzero"`
	Product_id int64  `json:",omitempty,omitzero"`
}
