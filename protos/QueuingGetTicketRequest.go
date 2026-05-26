package protos

type QueuingGetTicketRequest struct {
	RequestPacket
	NpSN           int64 `json:",omitempty,omitzero"`
	NpToken        string
	Npacode        string
	OSType         string
	AccessIP       string
	MakeStandby    bool `json:",omitempty,omitzero"`
	PassCheck      bool `json:",omitempty,omitzero"`
	PassCheckNexon bool `json:",omitempty,omitzero"`
	WaitingTicket  string
	ClientVersion  string
	NgsmToken      string
}
