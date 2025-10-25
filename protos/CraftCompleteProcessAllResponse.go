package protos

type CraftCompleteProcessAllResponse struct {
	ResponsePacket
	Protocol     Protocol
	CraftInfoDBs []CraftInfoDB
	TicketItemDB ItemDB
}
