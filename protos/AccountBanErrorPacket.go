package protos

type AccountBanErrorPacket struct {
	ResponsePacket
	Protocol  Protocol
	ErrorCode WebAPIErrorCode
	BanReason string
}
