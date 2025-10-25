package protos

type MomoTalkReadResponse struct {
	ResponsePacket
	Protocol          Protocol
	MomoTalkOutLineDB MomoTalkOutLineDB
	MomoTalkChoiceDBs []MomoTalkChoiceDB
}
