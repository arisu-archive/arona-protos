package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type OpenConditionDB struct {
	ContentType                flatdata.OpenConditionContent
	HideWhenLocked             bool
	AccountLevel               int64
	ScenarioModeId             int64
	CampaignStageUniqueId      int64
	MultipleConditionCheckType flatdata.MultipleConditionCheckType
	OpenDayOfWeek              flatdata.WeekDay
	OpenHour                   int64
	CloseDayOfWeek             flatdata.WeekDay
	CloseHour                  int64
	CafeIdForCafeRank          int64
	CafeRank                   int64
	OpenedCafeId               int64
}
