package protos

type MiniGameCCGEnterStageResponse struct {
	ResponsePacket
	Protocol Protocol
	StageDB  MiniGameCCGStagePlayDB
}
