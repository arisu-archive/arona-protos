package protos

type CharacterGearTierUpRequest struct {
	RequestPacket
	Protocol     Protocol
	GearServerId int64
	ReplaceInfos []SelectTicketReplaceInfo
}
