package protos

type AccountSetAdultCheckRequest struct {
	RequestPacket
	CheckAdultAgree bool `json:",omitempty,omitzero"`
}
