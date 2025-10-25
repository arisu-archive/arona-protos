package protos

type CharacterUnlockWeaponRequest struct {
	RequestPacket
	Protocol                Protocol
	TargetCharacterServerId int64
}
