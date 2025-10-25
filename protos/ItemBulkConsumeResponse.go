package protos

type ItemBulkConsumeResponse struct {
	ResponsePacket
	Protocol             Protocol
	UsedItemDB           ItemDB
	ParcelInfosInMailBox []ParcelInfo
}
