package protos

type CafeUpdatePresetFurnitureRequest struct {
	RequestPacket
	Protocol Protocol
	CafeDBId int64
	SlotId   int32
}
