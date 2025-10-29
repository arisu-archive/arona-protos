package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type OpenConditionListResponse struct {
	ResponsePacket
	ConditionContents []flatdata.OpenConditionContent `json:",omitempty,omitzero"`
}
