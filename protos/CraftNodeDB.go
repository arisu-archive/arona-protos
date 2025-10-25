package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CraftNodeDB struct {
	NodeTier           flatdata.CraftNodeTier
	SlotSequence       int64
	NodeId             int64
	NodeQuality        int64
	NodeLevel          int64
	NodeRandomSeed     int32
	NodeRandomSequence int32
	LeafNodeIds        []int64
	ResultId           int64
	CraftNodeResult    CraftNodeResult
}
