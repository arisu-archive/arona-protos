package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type PotentialGrowthRequestDB struct {
	Type flatdata.PotentialStatBonusRateType `json:",omitempty,omitzero"`
	Level int32 `json:",omitempty,omitzero"`
}
