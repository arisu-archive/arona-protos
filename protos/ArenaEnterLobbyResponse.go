package protos

import (
	"time"
)

type ArenaEnterLobbyResponse struct {
	ResponsePacket
	Protocol          Protocol
	ArenaPlayerInfoDB ArenaPlayerInfoDB
	OpponentUserDBs   []ArenaUserDB
	MapId             int64
	AutoRefreshTime   time.Time
}
