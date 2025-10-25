package protos

type MiniGameCCGSelectRewardCardResponse struct {
	ResponsePacket
	Protocol          Protocol
	StageDB           MiniGameCCGStagePlayDB
	SaveDB            MiniGameCCGSaveDB
	ReceivedRewardIds []int64
}
