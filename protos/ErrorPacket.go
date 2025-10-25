package protos

type ErrorPacket struct {
	ResponsePacket
	Protocol  Protocol
	Reason    string
	ErrorCode WebAPIErrorCode
}
