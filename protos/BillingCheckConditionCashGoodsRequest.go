package protos

type BillingCheckConditionCashGoodsRequest struct {
	RequestPacket
	User_id    string
	Product_id int64 `json:",omitempty,omitzero"`
}
