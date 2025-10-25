package protos

import (
	"time"
)

type ArenaSyncEchelonSettingTimeResponse struct {
	ResponsePacket
	Protocol           Protocol
	EchelonSettingTime time.Time
}
