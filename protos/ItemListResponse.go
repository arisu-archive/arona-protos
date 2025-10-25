package protos

type ItemListResponse struct {
	ResponsePacket
	Protocol      Protocol
	ItemDBs       []ItemDB
	ExpiryItemDBs []ItemDB
}
