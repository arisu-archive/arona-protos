package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CraftPresetNodeDB struct {
	NodeTier        flatdata.CraftNodeTier `json:",omitempty,omitzero"`
	IsActivated     bool                   `json:",omitempty,omitzero"`
	PriorityNodeIds []int64
	CostParcels     []*ParcelInfoImmutable
}
