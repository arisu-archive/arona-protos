package protos

type MiniGameCCGCreateGameRequest struct {
	RequestPacket
	Protocol         Protocol
	EventContentId   int64
	ForceDiscardSave bool
	DisablePerk      bool
}
