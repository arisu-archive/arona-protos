package protos

import (
	"time"
)

type MiniGameDreamMakerEndingDB struct {
	EventContentId int64
	EndingId       int64
	ClearCount     int64
	ClearDate      time.Time
}
