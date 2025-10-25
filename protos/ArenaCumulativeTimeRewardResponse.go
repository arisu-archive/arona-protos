package protos

import (
	"time"
)

type ArenaCumulativeTimeRewardResponse struct {
	ResponsePacket
	Protocol                 Protocol
	TimeRewardAmount         int64
	TimeRewardLastUpdateTime time.Time
	ParcelResult             ParcelResultDB
}
