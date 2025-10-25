package protos

type BillingCheckConditionCashGoodsRequest struct {
	RequestPacket
	Protocol   Protocol
	User_id    string
	Product_id int64
}
