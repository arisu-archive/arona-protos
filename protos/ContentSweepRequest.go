package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ContentSweepRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Content flatdata.ContentType `json:",omitempty,omitzero"`
	StageId int64 `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	Count int64 `json:",omitempty,omitzero"`
}
