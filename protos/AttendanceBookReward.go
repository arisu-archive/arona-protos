package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type AttendanceBookReward struct {
	UniqueId          int64
	Type              flatdata.AttendanceType
	AccountType       flatdata.AccountState
	DisplayOrder      int64
	AccountLevelLimit int64
	Title             string
	TitleImagePath    string
	CountRule         flatdata.AttendanceCountRule
	CountReset        flatdata.AttendanceResetType
	BookSize          int64
	StartDate         time.Time
	StartableEndDate  time.Time
	EndDate           time.Time
	ExpiryDate        int64
	MailType          flatdata.MailType
	DailyRewardIcons  map[int64]string
	DailyRewards      map[int64][]ParcelInfo
}
