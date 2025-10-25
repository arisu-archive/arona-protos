package protos

import (
	"time"
)

type ClanAssistRewardInfo struct {
	CharacterDBId            int64
	DeployDate               time.Time
	RentCount                int64
	CumultativeRewardParcels []ParcelInfo
	RentRewardParcels        []ParcelInfo
}
