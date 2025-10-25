package protos

type EventContentTreasureHistoryDB struct {
	EventContentId int64
	Round          int32
	Board          EventContentTreasureBoardHistory
	IsComplete     bool
	HintTreasures  []EventContentTreasureObject
}
