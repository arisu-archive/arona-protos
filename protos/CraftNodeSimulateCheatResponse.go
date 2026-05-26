package protos

import (
	"github.com/arisu-archive/mapx"
)

type CraftNodeSimulateCheatResponse struct {
	ResponsePacket
	NodeIdAndCount  *mapx.OrderedMap[int64, int32]
	SimulationCount int64 `json:",omitempty,omitzero"`
	Tier            int64 `json:",omitempty,omitzero"`
}
