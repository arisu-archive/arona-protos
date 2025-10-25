package protos

type CafeSummonCharacterRequest struct {
	RequestPacket
	Protocol          Protocol
	CafeDBId          int64
	CharacterServerId int64
}
