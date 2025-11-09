package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestManageBaseRequest struct {
	RequestPacket
	EventContentId int64                    `json:",omitempty,omitzero"`
	Difficulty     flatdata.StageDifficulty `json:",omitempty,omitzero"`
	TileUniqueId   int64                    `json:",omitempty,omitzero"`
	ManageCount    int32                    `json:",omitempty,omitzero"`
}
