package protos

type QueuingProcessWaitingQueueRequest struct {
	RequestPacket
	WaitingTicket string `json:",omitempty,omitzero"`
	ClientVersion string `json:",omitempty,omitzero"`
	OSType        string `json:",omitempty,omitzero"`
	AuthTicket    string `json:",omitempty,omitzero"`
	NpSN          int64  `json:",omitempty,omitzero"`
	NpToken       string `json:",omitempty,omitzero"`
	Npacode       string `json:",omitempty,omitzero"`
	AccessIP      string `json:",omitempty,omitzero"`
	NgsmToken     string `json:",omitempty,omitzero"`
}
