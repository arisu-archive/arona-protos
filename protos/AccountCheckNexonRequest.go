package protos

type AccountCheckNexonRequest struct {
	RequestPacket
	NpSN                 int64 `json:",omitempty,omitzero"`
	NpToken              string
	PassCheckNexonServer bool `json:",omitempty,omitzero"`
	EnterTicket          string
	ClientGeneratedKey   string
	ClientGeneratedIV    string
}
