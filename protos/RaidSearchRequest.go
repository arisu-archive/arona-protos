package protos

type RaidSearchRequest struct {
	RequestPacket
	Protocol   Protocol
	SecretCode string
	Tags       []string
}
