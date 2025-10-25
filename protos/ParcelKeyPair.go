package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ParcelKeyPair struct {
	Type flatdata.ParcelType
	Id   int64
}
