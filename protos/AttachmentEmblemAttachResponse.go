package protos

type AttachmentEmblemAttachResponse struct {
	ResponsePacket
	Protocol     Protocol
	AttachmentDB AccountAttachmentDB
}
