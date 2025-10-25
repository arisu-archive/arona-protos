package protos

type AccountCheckYostarResponse struct {
	ResponsePacket
	Protocol     Protocol
	ResultState  int32
	ResultMessag string
	Birth        string
	EncryptedKey string
	SignedKey    string
	EncryptedIV  string
	SignedIV     string
}
