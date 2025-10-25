package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidDB struct {
	Owner                         RaidMemberDescription
	ContentType                   flatdata.ContentType
	ServerId                      int64
	UniqueId                      int64
	SeasonId                      int64
	Begin                         time.Time
	End                           time.Time
	PlayerCount                   int64
	Tags                          []int32
	SecretCode                    string
	RaidState                     flatdata.RaidStatus
	IsPractice                    bool
	RaidBossDBs                   []RaidBossDB
	ParticipateCharacterServerIds map[int64][]int64
	IsEnterRoom                   bool
	AccountLevelWhenCreateDB      int64
	ClanAssistUsed                bool
}
