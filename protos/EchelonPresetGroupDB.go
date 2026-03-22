package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"github.com/arisu-archive/mapx"
)

type EchelonPresetGroupDB struct {
	GroupIndex    int32                                    `json:",omitempty,omitzero"`
	ExtensionType flatdata.EchelonExtensionType            `json:",omitempty,omitzero"`
	GroupLabel    string                                   `json:",omitempty,omitzero"`
	PresetDBs     *mapx.OrderedMap[int32, EchelonPresetDB] `json:",omitempty,omitzero"`
}
