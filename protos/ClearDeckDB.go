package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ClearDeckDB struct {
	ClearDeckCharacterDBs  []ClearDeckCharacterDB
	MulliganUniqueIds      []int64
	LeaderUniqueId         int64
	TSSInteractionUniqueId int64
	EchelonType            flatdata.EchelonType
	EchelonExtensionType   int64
}
