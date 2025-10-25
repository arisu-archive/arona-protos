package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EventContentMainGroundStageSaveDB struct {
	CampaignSubStageSaveDB
	ContentType flatdata.ContentType
}
