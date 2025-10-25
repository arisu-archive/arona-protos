package protos

type AttachmentGetResponse struct {
	ResponsePacket
	Protocol            Protocol
	AccountAttachmentDB AccountAttachmentDB
}
