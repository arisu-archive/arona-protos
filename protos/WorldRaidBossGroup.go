package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type WorldRaidBossGroup struct {
	ContentsValueChangeDB
	ContentsChangeType flatdata.ContentsChangeType
	GroupId            int64
	BossSpawnTime      time.Time
	EliminateTime      time.Time
}
