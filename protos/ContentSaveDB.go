package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ContentSaveDB struct {
	ContentType                 flatdata.ContentType
	AccountServerId             int64
	CreateTime                  time.Time
	StageUniqueId               int64
	LastEnterStageEchelonNumber int64
	StageEntranceFee            []ParcelInfo
	EnemyKillCountByUniqueId    map[int64]int64
	TacticClearTimeMscSum       int64
	AccountLevelWhenCreateDB    int64
	BIEchelon                   string
	BIEchelon1                  string
	BIEchelon2                  string
	BIEchelon3                  string
	BIEchelon4                  string
}
