package protos

type TBGEncounterDB struct {
	EncounterId                     int64
	InvokerServerId                 int64
	ObjectType                      int32
	ShouldDecreaseItemEffectCounter bool
	RewardUniqueIdByIndex           map[int32]int64
}
