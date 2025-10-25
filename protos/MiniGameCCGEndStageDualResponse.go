package protos

type MiniGameCCGEndStageDualResponse struct {
	ResponsePacket
	Protocol Protocol
	StageDB  MiniGameCCGStagePlayDB
	SaveDB   MiniGameCCGSaveDB
}
