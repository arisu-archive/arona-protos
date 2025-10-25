package protos

type CharacterWeaponExpGrowthRequest struct {
	RequestPacket
	Protocol                 Protocol
	TargetCharacterServerId  int64
	ConsumeUniqueIdAndCounts map[int64]int64
}
