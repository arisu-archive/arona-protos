package protos

type MailReceiveRequest struct {
	RequestPacket
	Protocol      Protocol
	MailServerIds []int64
}
