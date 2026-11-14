package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CharacterAllocateEngraveNodeRequest struct {
	RequestPacket
	TargetCharacterDBId int64                    `json:",omitempty,omitzero"`
	TreeType            flatdata.EngraveTreeType `json:",omitempty,omitzero"`
	NodeIndex           int32                    `json:",omitempty,omitzero"`
	ReplaceInfos        []*SelectTicketReplaceInfo
}
