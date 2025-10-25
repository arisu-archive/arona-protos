package protos

import (
	"time"
)

type ArenaOpponentListResponse struct {
	ResponsePacket
	Protocol        Protocol
	PlayerRank      int64
	OpponentUserDBs []ArenaUserDB
	AutoRefreshTime time.Time
}
