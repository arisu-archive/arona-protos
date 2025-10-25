package protos

type ShopBeforehandGachaSaveResponse struct {
	ResponsePacket
	Protocol            Protocol
	SelectGachaSnapshot BeforehandGachaSnapshotDB
}
