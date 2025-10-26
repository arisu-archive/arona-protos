package protos

type AccountVerifyAdultCheckResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CheckAdultAgree bool `json:",omitempty,omitzero"`
}
