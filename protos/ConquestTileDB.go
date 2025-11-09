package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestTileDB struct {
	EventContentId   int64                    `json:",omitempty,omitzero"`
	Difficulty       flatdata.StageDifficulty `json:",omitempty,omitzero"`
	TileUniqueId     int64                    `json:",omitempty,omitzero"`
	TileState        flatdata.TileState       `json:",omitempty,omitzero"`
	Level            int64                    `json:",omitempty,omitzero"`
	StarFlags        []bool                   `json:",omitempty,omitzero"`
	CreateTime       MxTime                   `json:",omitempty,omitzero"`
	IsThreeStarClear bool                     `json:",omitempty,omitzero"`
	IsAnyStarClear   bool                     `json:",omitempty,omitzero"`
}
