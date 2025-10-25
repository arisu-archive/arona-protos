package protos

type MailBoxFullErrorPacket struct {
	ResponsePacket
	Protocol  Protocol
	ErrorCode WebAPIErrorCode
}
