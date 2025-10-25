package protos

type EventContentTreasureObject struct {
	ServerId      int64
	RewardId      int64
	Rotation      int32
	IsHiddenImage bool
	Cells         []EventContentTreasureCell
}
