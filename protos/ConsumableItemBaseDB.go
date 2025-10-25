package protos

type ConsumableItemBaseDB struct {
	ParcelBase
	ServerId   int64
	UniqueId   int64
	StackCount int64
}
