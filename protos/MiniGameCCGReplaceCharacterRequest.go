package protos

type MiniGameCCGReplaceCharacterRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	SlotIndex      int32
	CharacterId    int64
	IsStriker      bool
}
