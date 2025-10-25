package protos

type NetworkTimeSyncResponse struct {
	ResponsePacket
	Protocol     Protocol
	ReceiveTick  int64
	EchoSendTick int64
}
