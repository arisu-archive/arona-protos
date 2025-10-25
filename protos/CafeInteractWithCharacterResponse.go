package protos

type CafeInteractWithCharacterResponse struct {
	ResponsePacket
	Protocol       Protocol
	CafeDB         CafeDB
	CharacterDB    CharacterDB
	ParcelResultDB ParcelResultDB
}
