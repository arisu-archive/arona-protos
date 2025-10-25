package protos

type QueuingGetTicketRequest struct {
	RequestPacket
	Protocol       Protocol
	NpSN           int64
	NpToken        string
	Npacode        string
	OSType         string
	AccessIP       string
	MakeStandby    bool
	PassCheck      bool
	PassCheckNexon bool
	WaitingTicket  string
	ClientVersion  string
	NgsmToken      string
}
