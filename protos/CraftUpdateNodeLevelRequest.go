package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CraftUpdateNodeLevelRequest struct {
	RequestPacket
	Protocol          Protocol
	ConsumeRequestDB  ConsumeRequestDB
	ConsumeGoldAmount int64
	SlotId            int64
	CraftNodeType     flatdata.CraftNodeTier
}
