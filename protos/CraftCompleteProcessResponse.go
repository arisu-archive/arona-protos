package protos

type CraftCompleteProcessResponse struct {
	ResponsePacket
	Protocol          Protocol
	AccountCurrencyDB AccountCurrencyDB
	CraftInfoDB       CraftInfoDB
	TicketItemDB      ItemDB
}
