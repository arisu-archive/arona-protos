package protos

type MiniGameCCGLobbyResponse struct {
	ResponsePacket
	Protocol    Protocol
	CCGSaveDB   MiniGameCCGSaveDB
	Perks       []int64
	RewardPoint int32
	CanSweep    bool
}
