package protos

import (
	"time"
)

type ShopRecruitDB struct {
	Id             int64
	SalesStartDate time.Time
	SalesEndDate   time.Time
	UpdateDate     time.Time
}
