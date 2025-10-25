package protos

type MailListResponse struct {
	ResponsePacket
	Protocol Protocol
	MailDBs  []MailDB
	Count    int64
}
