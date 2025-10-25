package protos

type CafeDeployFurnitureRequest struct {
	RequestPacket
	Protocol    Protocol
	CafeDBId    int64
	FurnitureDB FurnitureDB
}
