package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MailListRequest struct {
	RequestPacket
	IsReadMail      bool                     `json:",omitempty,omitzero"`
	PivotTime       MxTime                   `json:",omitempty,omitzero"`
	PivotIndex      int64                    `json:",omitempty,omitzero"`
	MailSortingRule flatdata.MailSortingRule `json:",omitempty,omitzero"`
	IsDescending    bool                     `json:",omitempty,omitzero"`
}
