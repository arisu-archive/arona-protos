package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidBossDB struct {
	ContentType     flatdata.ContentType
	BossIndex       int32
	BossCurrentHP   int64
	BossGroggyPoint int64
}
