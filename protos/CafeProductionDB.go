package protos

import (
	"time"
)

type CafeProductionDB struct {
	CafeDBId              int64
	ComfortValue          int64
	AppliedDate           time.Time
	ProductionParcelInfos []CafeProductionDB_CafeProductionParcelInfo
}
