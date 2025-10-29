package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ContentSaveDiscardRequest struct {
	RequestPacket
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
