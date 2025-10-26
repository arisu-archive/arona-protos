package protos

type AccountCheckNexonResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ResultState int32 `json:",omitempty,omitzero"`
	ResultMessage string `json:",omitempty,omitzero"`
	Birth string `json:",omitempty,omitzero"`
	EncryptedKey string `json:",omitempty,omitzero"`
	SignedKey string `json:",omitempty,omitzero"`
	EncryptedIV string `json:",omitempty,omitzero"`
	SignedIV string `json:",omitempty,omitzero"`
}
