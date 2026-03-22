package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"github.com/arisu-archive/mapx"
)

type NotificationEventContentReddotResponse struct {
	ResponsePacket
	Reddots                 *mapx.OrderedMap[int64, []flatdata.NotificationEventReddot]
	EventContentUnlockCGDBs *mapx.OrderedMap[int64, []EventContentCollectionDB]
}
