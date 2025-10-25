package protos

type AccountSetAdultCheckRequest struct {
	RequestPacket
	Protocol        Protocol
	CheckAdultAgree bool
}
