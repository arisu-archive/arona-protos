package protos

import (
	"time"
)

type ArenaDailyRewardResponse struct {
	ResponsePacket
	Protocol              Protocol
	ParcelResult          ParcelResultDB
	DailyRewardActiveTime time.Time
}
