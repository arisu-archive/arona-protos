package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestStageSaveDB struct {
	ContentSaveDB
	ContentType             flatdata.ContentType
	ConquestEventObjectDBId *int64
	EventContentId          int64
	Difficulty              flatdata.StageDifficulty
	TileUniqueId            int64
	TilePresetId            int64
	ConquestTileType        flatdata.ConquestTileType
	UseManageEchelon        bool
	AssistCharacterDB       AssistCharacterDB
	EchelonSlotType         int32
	EchelonSlotIndex        int32
}
