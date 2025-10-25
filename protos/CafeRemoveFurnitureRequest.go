package protos

type CafeRemoveFurnitureRequest struct {
	RequestPacket
	Protocol           Protocol
	CafeDBId           int64
	FurnitureServerIds []int64
}
