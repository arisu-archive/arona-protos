package protos

import (
	"time"
)

type ScenarioHistoryDB struct {
	ScenarioUniqueId int64
	ClearDateTime    time.Time
}
