package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MailListRequest struct {
	RequestPacket
	Protocol        Protocol
	IsReadMail      bool
	PivotTime       time.Time
	PivotIndex      int64
	MailSortingRule flatdata.MailSortingRule
	IsDescending    bool
}
