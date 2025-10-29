package protos

type AccountVerifyAdultCheckResponse struct {
	ResponsePacket
	CheckAdultAgree bool `json:",omitempty,omitzero"`
}
