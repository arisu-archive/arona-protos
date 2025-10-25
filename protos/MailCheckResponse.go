package protos

type MailCheckResponse struct {
	ResponsePacket
	Protocol Protocol
	Count    int64
}
