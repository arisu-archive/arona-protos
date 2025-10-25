package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ContentSaveDiscardRequest struct {
	RequestPacket
	Protocol      Protocol
	ContentType   flatdata.ContentType
	StageUniqueId int64
}
