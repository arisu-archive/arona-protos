package protos

type EventContentDiceResult struct {
	Index      int32
	MoveAmount int32
	Rewards    []ParcelInfo
}
