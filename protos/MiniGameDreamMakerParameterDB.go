package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MiniGameDreamMakerParameterDB struct {
	ParameterType flatdata.DreamMakerParameterType `json:",omitempty,omitzero"`
	BaseAmount int64 `json:",omitempty,omitzero"`
	CurrentAmount int64 `json:",omitempty,omitzero"`
}
