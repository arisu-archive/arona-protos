package protos

type MiniGameCCGBuyPerkRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	PerkId         int64
}
