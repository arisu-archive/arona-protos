package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ResetableContentId struct {
	Type   flatdata.ResetContentType `json:",omitempty,omitzero"`
	Mapped int64                     `json:",omitempty,omitzero"`
}
