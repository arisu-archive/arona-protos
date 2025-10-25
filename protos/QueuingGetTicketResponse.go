package protos

type QueuingGetTicketResponse struct {
	ResponsePacket
	Protocol               Protocol
	WaitingTicket          string
	EnterTicket            string
	TicketSequence         int64
	AllowedSequence        int64
	RequiredSecondsPerUser float64
	Birth                  string
	ServerSeed             string
}
