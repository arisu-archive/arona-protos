package protos

type CharacterExpGrowthRequest struct {
	RequestPacket
	Protocol                Protocol
	TargetCharacterServerId int64
	ConsumeRequestDB        ConsumeRequestDB
}
