package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type WeekDungeonStageHistoryDB struct {
	AccountServerId int64
	StageUniqueId   int64
	StarGoalRecord  map[flatdata.StarGoalType]int64
}
