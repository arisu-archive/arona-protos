package protos

type CharacterTranscendenceRequest struct {
	RequestPacket
	Protocol                Protocol
	TargetCharacterServerId int64
}
