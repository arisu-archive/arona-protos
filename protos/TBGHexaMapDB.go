package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type TBGHexaMapDB struct {
	MapType    flatdata.TBGThemaType
	Objects    map[int64]TBGHexaObjectDB
	IsTutorial bool
}
