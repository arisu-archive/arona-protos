package protos

type AccountDetachNexonRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
