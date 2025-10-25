package protos

type CharacterSkillLevelUpdateResponse struct {
	ResponsePacket
	Protocol       Protocol
	CharacterDB    CharacterDB
	ParcelResultDB ParcelResultDB
}
