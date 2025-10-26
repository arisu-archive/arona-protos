package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CampaignSubStageSaveDB struct {
	ContentSaveDB
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
}
