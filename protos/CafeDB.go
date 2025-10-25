package protos

import (
	"time"
)

type CafeDB struct {
	CafeDBId              int64
	CafeId                int64
	AccountId             int64
	CafeRank              int32
	LastUpdate            time.Time
	LastSummonDate        *time.Time
	IsNew                 bool
	CafeVisitCharacterDBs map[int64]CafeDB_CafeCharacterDB
	FurnitureDBs          []FurnitureDB
	ProductionAppliedTime time.Time
	ProductionDB          CafeProductionDB
}
