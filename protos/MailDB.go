package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type MailDB struct {
	ServerId          int64
	AccountServerId   int64
	Type              flatdata.MailType
	UniqueId          int64
	Sender            string
	LocalizedSender   map[flatdata.Language]string
	Comment           string
	LocalizedComment  map[flatdata.Language]string
	SendDate          time.Time
	ReceiptDate       *time.Time
	ExpireDate        *time.Time
	ParcelInfos       []ParcelInfo
	RemainParcelInfos []ParcelInfo
}
