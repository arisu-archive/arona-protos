package protos

type MomoTalkMessageListResponse struct {
	ResponsePacket
	Protocol          Protocol
	MomoTalkOutLineDB MomoTalkOutLineDB
	MomoTalkChoiceDBs []MomoTalkChoiceDB
}
