package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MissionNotifyUIEnteredRequest struct {
	RequestPacket
	UIType flatdata.MissionCompleteUIPrefabType `json:",omitempty,omitzero"`
}
