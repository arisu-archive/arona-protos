package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type TimeAttackDungeonBattleHistoryDB struct {
	DungeonType         flatdata.TimeAttackDungeonType
	GeasId              int64
	DefaultPoint        int64
	ClearTimePoint      int64
	EndFrame            int64
	MainCharacterDBs    []TimeAttackDungeonCharacterDB
	SupportCharacterDBs []TimeAttackDungeonCharacterDB
}
