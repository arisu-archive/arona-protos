package protos

type AttachmentEmblemAcquireRequest struct {
	RequestPacket
	Protocol  Protocol
	UniqueIds []int64
}
