package protos

type ItemSelectTicketRequest struct {
	RequestPacket
	Protocol           Protocol
	TicketItemServerId int64
	SelectItemUniqueId int64
	ConsumeCount       int32
}
