package protos

type ConsumeRequestDB struct {
	ConsumeItemServerIdAndCounts      map[int64]int64
	ConsumeEquipmentServerIdAndCounts map[int64]int64
	ConsumeFurnitureServerIdAndCounts map[int64]int64
}
