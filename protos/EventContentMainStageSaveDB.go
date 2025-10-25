package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EventContentMainStageSaveDB struct {
	CampaignMainStageSaveDB
	ContentType                flatdata.ContentType
	SelectedBuffDict           map[int64]int64
	CurrentAppearedBuffGroupId int64
}
