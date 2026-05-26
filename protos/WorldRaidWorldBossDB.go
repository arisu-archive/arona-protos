package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type WorldRaidWorldBossDB struct {
	ContentType  flatdata.ContentType
	GroupId      int64 `json:",omitempty,omitzero"`
	HP           int64 `json:",omitempty,omitzero"`
	Participants int64 `json:",omitempty,omitzero"`
}
