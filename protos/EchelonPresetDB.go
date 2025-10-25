package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EchelonPresetDB struct {
	GroupIndex             int32
	Index                  int32
	Label                  string
	LeaderUniqueId         int64
	TSSInteractionUniqueId int64
	StrikerUniqueIds       []int64
	SpecialUniqueIds       []int64
	CombatStyleIndex       []int32
	MulliganUniqueIds      []int64
	ExtensionType          flatdata.EchelonExtensionType
}
