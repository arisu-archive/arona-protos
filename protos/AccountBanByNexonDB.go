package protos

import (
	"time"
)

type AccountBanByNexonDB struct {
	AccountId    int64
	Npsn         int64
	AccountBanId int64
	BanType      int32
	BanDay       int32
	BanStartDate time.Time
	BanEndDate   time.Time
	BanComment   string
	DeleteFlag   int32
}
