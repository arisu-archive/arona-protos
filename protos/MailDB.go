package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"github.com/arisu-archive/mapx"
)

type MailDB struct {
	ServerId          int64             `json:",omitempty,omitzero"`
	AccountServerId   int64             `json:",omitempty,omitzero"`
	Type              flatdata.MailType `json:",omitempty,omitzero"`
	UniqueId          int64             `json:",omitempty,omitzero"`
	Sender            string
	LocalizedSender   *mapx.OrderedMap[string, string]
	Comment           string
	LocalizedComment  *mapx.OrderedMap[string, string]
	SendDate          MxTime  `json:",omitempty,omitzero"`
	ReceiptDate       *MxTime `json:",omitempty,omitzero"`
	ExpireDate        *MxTime `json:",omitempty,omitzero"`
	ParcelInfos       []ParcelInfo
	RemainParcelInfos []ParcelInfo
}
