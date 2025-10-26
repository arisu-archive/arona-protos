package protos

type AccountDetachNexonResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ResultState int32 `json:",omitempty,omitzero"`
	ResultMessage string `json:",omitempty,omitzero"`
}
