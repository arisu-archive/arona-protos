package protos

type ContentSweepMultiSweepRequest struct {
	RequestPacket
	Protocol             Protocol
	MultiSweepParameters []MultiSweepParameter
}
