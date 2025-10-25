package protos

type InventoryFullErrorPacket struct {
	ResponsePacket
	Protocol    Protocol
	ErrorCode   WebAPIErrorCode
	ParcelInfos []ParcelInfo
}
