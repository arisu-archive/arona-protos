package protos

type CafeApplyTemplateRequest struct {
	RequestPacket
	Protocol              Protocol
	TemplateId            int64
	CafeDBId              int64
	UseOtherCafeFurniture bool
}
