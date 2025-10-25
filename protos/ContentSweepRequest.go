package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ContentSweepRequest struct {
	RequestPacket
	Protocol       Protocol
	Content        flatdata.ContentType
	StageId        int64
	EventContentId int64
	Count          int64
}
