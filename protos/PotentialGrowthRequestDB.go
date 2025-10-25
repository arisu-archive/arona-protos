package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type PotentialGrowthRequestDB struct {
	Type  flatdata.PotentialStatBonusRateType
	Level int32
}
