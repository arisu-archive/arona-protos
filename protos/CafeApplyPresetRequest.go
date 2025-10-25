package protos

type CafeApplyPresetRequest struct {
	RequestPacket
	Protocol              Protocol
	SlotId                int32
	CafeDBId              int64
	UseOtherCafeFurniture bool
}
