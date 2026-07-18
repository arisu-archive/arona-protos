package protos

import (
	"github.com/arisu-archive/mapx"
)

type CraftSimulateCheatResponse struct {
	ResponsePacket
	ParcelIdAndCount *mapx.OrderedMap[int64, int32]
	SimulationCount  int64 `json:",omitempty,omitzero"`
	NodeId           int64 `json:",omitempty,omitzero"`
	Message          string
}
