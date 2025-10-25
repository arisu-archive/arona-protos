package protos

import (
	"time"
)

type EventContentCollectionDB struct {
	EventContentId int64
	GroupId        int64
	UniqueId       int64
	ReceiveDate    time.Time
}
