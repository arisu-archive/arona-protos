package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestMainStorySummary struct {
	EventContentId          int64                                  `json:",omitempty,omitzero"`
	Difficulty              flatdata.StageDifficulty               `json:",omitempty,omitzero"`
	ConquestStepSummaryDict map[int32]ConquestMainStoryStepSummary `json:",omitempty,omitzero"`
}
