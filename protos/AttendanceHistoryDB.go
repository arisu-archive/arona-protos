package protos

import (
	"time"
)

type AttendanceHistoryDB struct {
	ServerId               int64
	AttendanceBookUniqueId int64
	AttendedDay            map[int64]time.Time
	Expired                bool
}
