package protos

type EquipmentItemTierUpRequest struct {
	RequestPacket
	Protocol                Protocol
	TargetEquipmentServerId int64
	ReplaceInfos            []SelectTicketReplaceInfo
}
