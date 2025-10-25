package protos

type AccountVerifyAdultCheckResponse struct {
	ResponsePacket
	Protocol        Protocol
	CheckAdultAgree bool
}
