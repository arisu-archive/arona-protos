package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ParcelKeyPair struct {
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	Id int64 `json:",omitempty,omitzero"`
}
