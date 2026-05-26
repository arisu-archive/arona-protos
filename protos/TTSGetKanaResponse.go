package protos

type TTSGetKanaResponse struct {
	ResponsePacket
	CallName             string
	ActualCallName       string
	CallNameKatakana     string
	CallNameKorean       string
	ActualCallNameKorean string
}
