package protos

type EquipmentItemEquipResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
