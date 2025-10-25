package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CraftPresetNodeDB struct {
	NodeTier         flatdata.CraftNodeTier
	IsActivated      bool
	PriortyNodeId    int64
	ConsumeRequestDB ConsumeRequestDB
}
