package protos

type QueuingGetTicketRequest struct {
	RequestPacket
	NpSN           int64  `json:",omitempty,omitzero"`
	NpToken        string `json:",omitempty,omitzero"`
	Npacode        string `json:",omitempty,omitzero"`
	OSType         string `json:",omitempty,omitzero"`
	AccessIP       string `json:",omitempty,omitzero"`
	MakeStandby    bool   `json:",omitempty,omitzero"`
	PassCheck      bool   `json:",omitempty,omitzero"`
	PassCheckNexon bool   `json:",omitempty,omitzero"`
	WaitingTicket  string `json:",omitempty,omitzero"`
	ClientVersion  string `json:",omitempty,omitzero"`
	NgsmToken      string `json:",omitempty,omitzero"`
}
