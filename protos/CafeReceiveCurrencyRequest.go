package protos

type CafeReceiveCurrencyRequest struct {
	RequestPacket
	Protocol        Protocol
	AccountServerId int64
	CafeDBId        int64
}
