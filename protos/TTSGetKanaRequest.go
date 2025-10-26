package protos

type TTSGetKanaRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CallName string `json:",omitempty,omitzero"`
}
