package protos

type AccountCheckYostarRequest struct {
	RequestPacket
	Protocol         Protocol
	UID              int64
	YostarToken      string
	EnterTicket      string
	PassCookieResult bool
	Cookie           string
}
