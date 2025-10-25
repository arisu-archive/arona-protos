package protos

type AccountCheckNexonResponse struct {
	ResponsePacket
	Protocol      Protocol
	ResultState   int32
	ResultMessage string
	Birth         string
	EncryptedKey  string
	SignedKey     string
	EncryptedIV   string
	SignedIV      string
}
