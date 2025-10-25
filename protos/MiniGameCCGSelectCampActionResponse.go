package protos

type MiniGameCCGSelectCampActionResponse struct {
	ResponsePacket
	Protocol Protocol
	StageDB  MiniGameCCGStagePlayDB
	SaveDB   MiniGameCCGSaveDB
}
