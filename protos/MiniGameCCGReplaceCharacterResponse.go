package protos

type MiniGameCCGReplaceCharacterResponse struct {
	ResponsePacket
	Protocol       Protocol
	SaveDB         MiniGameCCGSaveDB
	CCGCharacterDB MiniGameCCGCharacterDB
}
