package protos

type ShopBeforehandGachaRunResponse struct {
	ResponsePacket
	Protocol            Protocol
	SelectGachaSnapshot BeforehandGachaSnapshotDB
}
