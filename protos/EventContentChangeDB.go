package protos

import (
	"time"
)

type EventContentChangeDB struct {
	EventContentId        int64
	UseAmount             int64
	ChangeCount           int64
	AccumulateChangeCount int64
	LastUpdateDate        time.Time
	ChangeFlag            bool
}
