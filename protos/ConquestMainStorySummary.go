package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestMainStorySummary struct {
	EventContentId          int64
	Difficulty              flatdata.StageDifficulty
	ConquestStepSummaryDict map[int32]ConquestMainStoryStepSummary
}
