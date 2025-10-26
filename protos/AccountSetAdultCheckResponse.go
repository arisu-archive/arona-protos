package protos

type AccountSetAdultCheckResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
