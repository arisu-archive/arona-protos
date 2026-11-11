package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"github.com/arisu-archive/mapx"
)

type TBGHexaMapDB struct {
	MapType    flatdata.TBGThemaType                     `json:",omitempty,omitzero"`
	Objects    *mapx.OrderedMap[int64, *TBGHexaObjectDB] `json:"obj"`
	IsTutorial bool                                      `json:"tutorial,omitempty,omitzero"`
}
