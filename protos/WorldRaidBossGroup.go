package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type WorldRaidBossGroup struct {
	ContentsValueChangeDB
	ContentType   flatdata.ContentType
	GroupId       int64  `json:",omitempty,omitzero"`
	BossSpawnTime MxTime `json:",omitempty,omitzero"`
	EliminateTime MxTime `json:",omitempty,omitzero"`
}
