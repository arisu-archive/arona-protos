package protos

type ToastListResponse struct {
	ResponsePacket
	Protocol Protocol
	ToastDBs []ToastDB
}
