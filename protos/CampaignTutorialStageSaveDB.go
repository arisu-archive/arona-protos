package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CampaignTutorialStageSaveDB struct {
	ContentSaveDB
	ContentType flatdata.ContentType
}
