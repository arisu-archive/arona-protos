package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type OpenConditionListResponse struct {
	ResponsePacket
	Protocol          Protocol
	ConditionContents []flatdata.OpenConditionContent
}
