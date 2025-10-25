package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type StatSnapshot struct {
	Stat  flatdata.StatType
	Start int64
	End   int64
}
