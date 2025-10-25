package protos

type AccountCheckNexonRequest struct {
	RequestPacket
	Protocol             Protocol
	NpSN                 int64
	NpToken              string
	PassCheckNexonServer bool
	EnterTicket          string
	ClientGeneratedKey   string
	ClientGeneratedIV    string
}
