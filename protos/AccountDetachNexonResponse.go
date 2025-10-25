package protos

type AccountDetachNexonResponse struct {
	ResponsePacket
	Protocol      Protocol
	ResultState   int32
	ResultMessage string
}
