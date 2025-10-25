package protos

type ClanMyAssistListResponse struct {
	ResponsePacket
	Protocol          Protocol
	ClanAssistSlotDBs []ClanAssistSlotDB
}
