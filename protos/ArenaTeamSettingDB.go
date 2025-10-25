package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ArenaTeamSettingDB struct {
	EchelonType               flatdata.EchelonType
	LeaderCharacterId         int64
	TSSInteractionCharacterId int64
	MainCharacters            []ArenaCharacterDB
	SupportCharacters         []ArenaCharacterDB
	TSSCharacterDB            ArenaCharacterDB
	MapId                     int64
}
