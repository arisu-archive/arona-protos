package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MemoryLobbyDB struct {
	ParcelBase
	Type                flatdata.ParcelType
	MemoryLobbyUniqueId int64
}
