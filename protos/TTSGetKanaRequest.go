package protos

type TTSGetKanaRequest struct {
	RequestPacket
	CallName string `json:",omitempty,omitzero"`
}
