package protos

type EchelonSaveRequest struct {
	RequestPacket
	Protocol       Protocol
	EchelonDB      EchelonDB
	AssistUseInfos []ClanAssistUseInfo
	IsPractice     bool
}
