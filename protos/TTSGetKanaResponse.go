package protos

type TTSGetKanaResponse struct {
	ResponsePacket
	Protocol             Protocol
	CallName             string
	ActualCallName       string
	CallNameKatakana     string
	CallNameKorean       string
	ActualCallNameKorean string
}
