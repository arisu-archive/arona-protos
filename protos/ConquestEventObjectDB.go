package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestEventObjectDB struct {
	ConquestObjectDBId int64
	EventContentId     int64
	Difficulty         flatdata.StageDifficulty
	TileUniqueId       int64
	ObjectId           int64
	ObjectType         ConquestEventObjectType
	IsAlive            bool
}
