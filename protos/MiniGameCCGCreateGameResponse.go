package protos

type MiniGameCCGCreateGameResponse struct {
	ResponsePacket
	Protocol  Protocol
	CCGSaveDB MiniGameCCGSaveDB
}
