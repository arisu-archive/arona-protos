package protos

type ItemBulkConsumeRequest struct {
	RequestPacket
	Protocol           Protocol
	TargetItemServerId int64
	ConsumeCount       int32
}
