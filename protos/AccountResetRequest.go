package protos

type AccountResetRequest struct {
	RequestPacket
	Protocol Protocol
	DevId    string
}
