package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClearDeckKey struct {
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
	Arguments []int64 `json:",omitempty,omitzero"`
}
