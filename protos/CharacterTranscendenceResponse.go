package protos

type CharacterTranscendenceResponse struct {
	ResponsePacket
	Protocol       Protocol
	CharacterDB    CharacterDB
	ParcelResultDB ParcelResultDB
}
