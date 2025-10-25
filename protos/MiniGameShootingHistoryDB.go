package protos

import (
	"time"
)

type MiniGameShootingHistoryDB struct {
	EventContentId int64
	UniqueId       int64
	ArriveSection  int64
	LastUpdateDate time.Time
	IsClearToday   bool
}
