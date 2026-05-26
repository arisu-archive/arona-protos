package protos

type AccountCheckNexonResponse struct {
	ResponsePacket
	ResultState   int32 `json:",omitempty,omitzero"`
	ResultMessage string
	Birth         string
	EncryptedKey  string
	SignedKey     string
	EncryptedIV   string
	SignedIV      string
}
