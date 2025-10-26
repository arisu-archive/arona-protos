package protos

type AccountVerifyAdultCheckRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
