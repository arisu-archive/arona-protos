package protos

type CampaignEnterMainStageStrategySkipRequest struct {
	RequestPacket
	Protocol                    Protocol
	StageUniqueId               int64
	LastEnterStageEchelonNumber int64
}
