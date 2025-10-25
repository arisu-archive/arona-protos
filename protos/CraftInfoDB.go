package protos

import (
	"time"
)

type CraftInfoDB struct {
	SlotSequence      int64
	StartTime         time.Time
	EndTime           time.Time
	CraftSlotOpenDate time.Time
	Nodes             []CraftNodeDB
}
