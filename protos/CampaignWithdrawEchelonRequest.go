package protos

type CampaignWithdrawEchelonRequest struct {
	RequestPacket
	Protocol                Protocol
	StageUniqueId           int64
	WithdrawEchelonEntityId []int64
}
