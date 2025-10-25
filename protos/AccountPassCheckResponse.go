package protos

type AccountPassCheckResponse struct {
	ResponsePacket
	Protocol     Protocol
	EncryptedKey string
	SignedKey    string
	EncryptedIV  string
	SignedIV     string
}
