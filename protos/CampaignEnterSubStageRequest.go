package protos

type CampaignEnterSubStageRequest struct {
	RequestPacket
	Protocol                    Protocol
	StageUniqueId               int64
	LastEnterStageEchelonNumber int64
}
