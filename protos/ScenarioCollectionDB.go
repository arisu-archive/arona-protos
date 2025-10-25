package protos

import (
	"time"
)

type ScenarioCollectionDB struct {
	GroupId     int64
	UniqueId    int64
	ReceiveDate time.Time
}
