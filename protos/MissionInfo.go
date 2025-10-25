package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MissionInfo struct {
	IMissionConstraint
	Id                            int64
	Category                      flatdata.MissionCategory
	ResetType                     flatdata.MissionResetType
	ToastDisplayType              flatdata.MissionToastDisplayConditionType
	Description                   uint32
	IsVisible                     bool
	IsLimited                     bool
	StartDate                     time.Time
	StartableEndDate              time.Time
	EndDate                       time.Time
	EndDday                       int64
	AccountState                  flatdata.AccountState
	AccountLevel                  int64
	PreMissionIds                 []int64
	NextMissionId                 int64
	SuddenMissionContentTypes     []flatdata.SuddenMissionContentType
	CompleteConditionType         flatdata.MissionCompleteConditionType
	CompleteConditionCount        int64
	CompleteConditionParameters   []int64
	Tags                          []flatdata.Tag
	CompleteConditionMissionIds   []int64
	CompleteConditionMissionCount int64
	CompleteConditionRewards      []ParcelInfo
	RewardIcon                    string
	Rewards                       []ParcelInfo
	DateAutoRefer                 flatdata.ContentType
	ToastImagePath                string
	DisplayOrder                  int64
	HasFollowingMission           bool
	Shortcuts                     []string
	ChallengeStageId              int64
}
