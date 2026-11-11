package protos

type TTSGetKanaResponse struct {
	ResponsePacket
	CallName             string `json:",omitempty,omitzero"`
	ActualCallName       string `json:",omitempty,omitzero"`
	CallNameKatakana     string `json:",omitempty,omitzero"`
	CallNameKorean       string `json:",omitempty,omitzero"`
	ActualCallNameKorean string `json:",omitempty,omitzero"`
}
