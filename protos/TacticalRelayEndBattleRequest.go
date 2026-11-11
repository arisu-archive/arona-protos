package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type TacticalRelayEndBattleRequest struct {
	RequestPacket
	SeasonId           int64                    `json:",omitempty,omitzero"`
	StageId            int64                    `json:",omitempty,omitzero"`
	Summary            *BattleSummary           `json:",omitempty,omitzero"`
	SelectedRewardType flatdata.EngraveTreeType `json:",omitempty,omitzero"`
}
