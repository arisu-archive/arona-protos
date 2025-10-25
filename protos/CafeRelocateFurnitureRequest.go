package protos

type CafeRelocateFurnitureRequest struct {
	RequestPacket
	Protocol    Protocol
	CafeDBId    int64
	FurnitureDB FurnitureDB
}
