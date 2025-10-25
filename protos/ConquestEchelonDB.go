package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestEchelonDB struct {
	EventContentId          int64
	Difficulty              flatdata.StageDifficulty
	TileUniqueId            int64
	EchelonDB               EchelonDB
	AssistCharacterUniqueId int64
	AssistUseInfo           ClanAssistUseInfo
}
