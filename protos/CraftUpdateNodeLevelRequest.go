package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CraftUpdateNodeLevelRequest struct {
	RequestPacket
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
	ConsumeGoldAmount int64 `json:",omitempty,omitzero"`
	SlotId int64 `json:",omitempty,omitzero"`
	CraftNodeType flatdata.CraftNodeTier `json:",omitempty,omitzero"`
}
