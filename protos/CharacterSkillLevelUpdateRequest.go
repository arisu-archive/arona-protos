package protos

type CharacterSkillLevelUpdateRequest struct {
	RequestPacket
	Protocol            Protocol
	TargetCharacterDBId int64
	SkillSlot           SkillSlot
	Level               int32
	ReplaceInfos        []SelectTicketReplaceInfo
}
