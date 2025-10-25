package protos

type CharacterPotentialGrowthRequest struct {
	RequestPacket
	Protocol                  Protocol
	TargetCharacterDBId       int64
	PotentialGrowthRequestDBs []PotentialGrowthRequestDB
	ReplaceInfos              []SelectTicketReplaceInfo
}
