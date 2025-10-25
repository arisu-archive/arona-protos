package protos

type EventContentWithdrawEchelonRequest struct {
	RequestPacket
	Protocol                Protocol
	EventContentId          int64
	StageUniqueId           int64
	WithdrawEchelonEntityId []int64
}
