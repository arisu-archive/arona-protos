package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type AccountContentBanErrorPacket struct {
	ResponsePacket
	ErrorCode    WebAPIErrorCode `json:",omitempty,omitzero"`
	ContentType  flatdata.ContentType
	BanStartDate MxTime `json:",omitempty,omitzero"`
	BanEndDate   MxTime `json:",omitempty,omitzero"`
}
