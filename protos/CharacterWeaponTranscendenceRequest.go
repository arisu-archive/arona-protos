package protos

type CharacterWeaponTranscendenceRequest struct {
	RequestPacket
	Protocol                Protocol
	TargetCharacterServerId int64
}
