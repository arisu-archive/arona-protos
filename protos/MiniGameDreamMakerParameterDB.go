package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MiniGameDreamMakerParameterDB struct {
	ParameterType flatdata.DreamMakerParameterType
	BaseAmount    int64
	CurrentAmount int64
}
