package protos

type AccountPassCheckRequest struct {
	RequestPacket
	Protocol           Protocol
	DevId              string
	OnlyAccountId      bool
	ClientGeneratedKey string
	ClientGeneratedIV  string
}
