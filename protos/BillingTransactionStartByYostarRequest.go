package protos

type BillingTransactionStartByYostarRequest struct {
	RequestPacket
	Protocol                    Protocol
	ShopCashId                  int64
	ShopCashProductSelectionDBs []ShopCashProductSelectionDB
	VirtualPayment              bool
}
