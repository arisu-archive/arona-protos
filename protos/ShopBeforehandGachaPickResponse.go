package protos

type ShopBeforehandGachaPickResponse struct {
	ResponsePacket
	Protocol      Protocol
	GachaResults  []GachaResult
	AcquiredItems []ItemDB
}
