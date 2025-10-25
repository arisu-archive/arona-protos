package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ResetableContentId struct {
	Type   flatdata.ResetContentType
	Mapped int64
}
