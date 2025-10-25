package protos

type CafeOpenResponse struct {
	ResponsePacket
	Protocol     Protocol
	OpenedCafeDB CafeDB
	FurnitureDBs []FurnitureDB
}
