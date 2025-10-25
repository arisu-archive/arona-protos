package protos

type AttachmentEmblemAcquireResponse struct {
	ResponsePacket
	Protocol  Protocol
	EmblemDBs []EmblemDB
}
