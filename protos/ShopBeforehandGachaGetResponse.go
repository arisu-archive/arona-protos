package protos

type ShopBeforehandGachaGetResponse struct {
	ResponsePacket
	Protocol                Protocol
	AlreadyPicked           bool
	BeforehandGachaSnapshot BeforehandGachaSnapshotDB
}
