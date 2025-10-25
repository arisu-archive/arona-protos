package protos

type CharacterGearUnlockRequest struct {
	RequestPacket
	Protocol          Protocol
	CharacterServerId int64
	SlotIndex         int32
}
