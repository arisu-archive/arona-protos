package protos

type ItemSelectTicketResponse struct {
	ResponsePacket
	Protocol       Protocol
	UsedItemDB     ItemDB
	ParcelResultDB ParcelResultDB
}
