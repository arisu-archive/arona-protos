package protos

type CharacterFavorGrowthResponse struct {
	ResponsePacket
	Protocol                     Protocol
	CharacterDB                  CharacterDB
	ConsumeStackableItemDBResult []ItemDB
	ParcelResultDB               ParcelResultDB
}
