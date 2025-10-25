package protos

type ItemConsumeRequest struct {
	RequestPacket
	Protocol           Protocol
	TargetItemServerId int64
	ConsumeCount       int32
}
