package protos

type AccountCallNameRequest struct {
	RequestPacket
	CallName string `json:",omitempty,omitzero"`
	CallNameKatakana string `json:",omitempty,omitzero"`
	CallNameKorean string `json:",omitempty,omitzero"`
}
