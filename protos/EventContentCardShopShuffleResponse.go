package protos

type EventContentCardShopShuffleResponse struct {
	ResponsePacket
	Protocol           Protocol
	CardShopElementDBs []CardShopElementDB
}
