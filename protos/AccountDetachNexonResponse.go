package protos

type AccountDetachNexonResponse struct {
	ResponsePacket
	ResultState int32 `json:",omitempty,omitzero"`
	ResultMessage string `json:",omitempty,omitzero"`
}
