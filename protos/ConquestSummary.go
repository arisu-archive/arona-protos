package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"github.com/arisu-archive/mapx"
)

type ConquestSummary struct {
	EventContentId          int64                    `json:",omitempty,omitzero"`
	Difficulty              flatdata.StageDifficulty `json:",omitempty,omitzero"`
	ConquestStepSummaryDict *mapx.OrderedMap[int32, ConquestStepSummary]
}
