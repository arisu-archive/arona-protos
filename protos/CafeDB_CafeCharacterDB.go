package protos

import (
	"time"
)

type CafeDB_CafeCharacterDB struct {
	VisitingCharacterDB
	IsSummon         bool
	LastInteractTime time.Time
}
