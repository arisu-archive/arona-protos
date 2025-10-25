package protos

type AttachmentEmblemListResponse struct {
	ResponsePacket
	Protocol  Protocol
	EmblemDBs []EmblemDB
}
