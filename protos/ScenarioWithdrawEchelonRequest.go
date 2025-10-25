package protos

type ScenarioWithdrawEchelonRequest struct {
	RequestPacket
	Protocol                Protocol
	StageUniqueId           int64
	WithdrawEchelonEntityId []int64
}
