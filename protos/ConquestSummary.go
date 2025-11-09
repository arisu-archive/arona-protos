package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestSummary struct {
	EventContentId          int64                         `json:",omitempty,omitzero"`
	Difficulty              flatdata.StageDifficulty      `json:",omitempty,omitzero"`
	ConquestStepSummaryDict map[int32]ConquestStepSummary `json:",omitempty,omitzero"`
}
