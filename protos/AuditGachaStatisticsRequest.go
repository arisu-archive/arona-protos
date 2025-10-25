package protos

type AuditGachaStatisticsRequest struct {
	RequestPacket
	Protocol            Protocol
	MerchandiseUniqueId int64
	ShopUniqueId        int64
	Count               int64
}
