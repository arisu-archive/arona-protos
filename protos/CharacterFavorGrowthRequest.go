package protos

type CharacterFavorGrowthRequest struct {
	RequestPacket
	Protocol                  Protocol
	TargetCharacterDBId       int64
	ConsumeItemDBIdsAndCounts map[int64]int32
}
