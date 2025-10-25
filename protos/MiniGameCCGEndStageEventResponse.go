package protos

type MiniGameCCGEndStageEventResponse struct {
	ResponsePacket
	Protocol Protocol
	StageDB  MiniGameCCGStagePlayDB
	SaveDB   MiniGameCCGSaveDB
}
