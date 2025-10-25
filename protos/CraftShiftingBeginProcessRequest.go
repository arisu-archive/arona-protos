package protos

type CraftShiftingBeginProcessRequest struct {
	RequestPacket
	Protocol         Protocol
	SlotId           int64
	RecipeId         int64
	ConsumeRequestDB ConsumeRequestDB
}
