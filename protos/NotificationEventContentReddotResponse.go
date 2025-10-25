package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type NotificationEventContentReddotResponse struct {
	ResponsePacket
	Protocol                Protocol
	Reddots                 map[int64][]flatdata.NotificationEventReddot
	EventContentUnlockCGDBs map[int64][]EventContentCollectionDB
}
