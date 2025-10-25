package protos

type ItemSellRequest struct {
	RequestPacket
	Protocol        Protocol
	TargetServerIds []int64
}
