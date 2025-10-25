package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidTeamSettingDB struct {
	AccountId                     int64
	TryNumber                     int64
	EchelonType                   flatdata.EchelonType
	EchelonExtensionType          flatdata.EchelonExtensionType
	MainCharacterDBs              []RaidCharacterDB
	SupportCharacterDBs           []RaidCharacterDB
	SkillCardMulliganCharacterIds []int64
	TSSInteractionUniqueId        int64
	LeaderCharacterUniqueId       int64
}
