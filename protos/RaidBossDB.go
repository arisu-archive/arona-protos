package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidBossDB struct {
	ContentType     flatdata.ContentType
	BossIndex       int32 `json:",omitempty,omitzero"`
	BossCurrentHP   int64 `json:",omitempty,omitzero"`
	BossGroggyPoint int64 `json:",omitempty,omitzero"`
}
