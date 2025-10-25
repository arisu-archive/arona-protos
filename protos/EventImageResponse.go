package protos

type EventImageResponse struct {
	ResponsePacket
	Protocol   Protocol
	ImageBytes []byte
}
