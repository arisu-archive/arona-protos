package protos

type AccountDismissRepurchasablePopupRequest struct {
	RequestPacket
	Protocol   Protocol
	ProductIds []int64
}
