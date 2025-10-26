package protos

type AccountCheckNexonRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	NpSN int64 `json:",omitempty,omitzero"`
	NpToken string `json:",omitempty,omitzero"`
	PassCheckNexonServer bool `json:",omitempty,omitzero"`
	EnterTicket string `json:",omitempty,omitzero"`
	ClientGeneratedKey string `json:",omitempty,omitzero"`
	ClientGeneratedIV string `json:",omitempty,omitzero"`
}
