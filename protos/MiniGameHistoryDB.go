package protos

import (
	"time"
)

type MiniGameHistoryDB struct {
	EventContentId   int64
	UniqueId         int64
	HighScore        int64
	AccumulatedScore int64
	ClearDate        time.Time
	IsFullCombo      bool
}
