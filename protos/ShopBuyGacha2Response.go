package protos

import (
	"time"
)

type ShopBuyGacha2Response struct {
	ResponsePacket
	Protocol       Protocol
	UpdateTime     time.Time
	GemBonusRemain int64
	GemPaidRemain  int64
	ConsumedItems  []ItemDB
	GachaResults   []GachaResult
	AcquiredItems  []ItemDB
}
