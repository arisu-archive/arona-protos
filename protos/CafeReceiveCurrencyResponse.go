package protos

type CafeReceiveCurrencyResponse struct {
	ResponsePacket
	Protocol       Protocol
	CafeDB         CafeDB
	CafeDBs        []CafeDB
	ParcelResultDB ParcelResultDB
}
