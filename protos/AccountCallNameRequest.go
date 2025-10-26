package protos

type AccountCallNameRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CallName string `json:",omitempty,omitzero"`
	CallNameKatakana string `json:",omitempty,omitzero"`
	CallNameKorean string `json:",omitempty,omitzero"`
}
