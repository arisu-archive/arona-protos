package protos

type AccountSetAdultCheckRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CheckAdultAgree bool `json:",omitempty,omitzero"`
}
