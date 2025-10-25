package protos

type EventContentCardShopListResponse struct {
	ResponsePacket
	Protocol           Protocol
	CardShopElementDBs []CardShopElementDB
	RewardHistory      map[int64][]ParcelInfo
}
