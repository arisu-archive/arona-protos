package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EchelonPresetGroupDB struct {
	GroupIndex    int32
	ExtensionType flatdata.EchelonExtensionType
	GroupLabel    string
	PresetDBs     map[int32]EchelonPresetDB
}
