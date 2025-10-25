package protos

type CafeInteractWithCharacterRequest struct {
	RequestPacket
	Protocol    Protocol
	CafeDBId    int64
	CharacterId int64
}
