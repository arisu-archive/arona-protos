package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClearDeckKey struct {
	ContentType flatdata.ContentType
	Arguments   []int64
}
