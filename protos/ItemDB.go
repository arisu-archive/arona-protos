package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ItemDB struct {
	ConsumableItemBaseDB
	Type flatdata.ParcelType
}
