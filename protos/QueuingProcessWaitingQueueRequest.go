package protos

type QueuingProcessWaitingQueueRequest struct {
	RequestPacket
	WaitingTicket string
	ClientVersion string
	OSType        string
	AuthTicket    string
	NpSN          int64 `json:",omitempty,omitzero"`
	NpToken       string
	Npacode       string
	AccessIP      string
	NgsmToken     string
}
