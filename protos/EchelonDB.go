package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EchelonDB struct {
	AccountServerId               int64
	EchelonType                   flatdata.EchelonType
	EchelonNumber                 int64
	ExtensionType                 flatdata.EchelonExtensionType
	LeaderServerId                int64
	MainSlotServerIds             []int64
	SupportSlotServerIds          []int64
	TSSInteractionServerId        int64
	UsingFlag                     EchelonDB_EchelonStatusFlag
	SkillCardMulliganCharacterIds []int64
	CombatStyleIndex              []int32
}
