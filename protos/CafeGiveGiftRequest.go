package protos

type CafeGiveGiftRequest struct {
	RequestPacket
	Protocol          Protocol
	CafeDBId          int64
	CharacterUniqueId int64
	ConsumeRequestDB  ConsumeRequestDB
}
