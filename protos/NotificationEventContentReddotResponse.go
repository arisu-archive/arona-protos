package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type NotificationEventContentReddotResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Reddots map[int64][]flatdata.NotificationEventReddot `json:",omitempty,omitzero"`
	EventContentUnlockCGDBs map[int64][]EventContentCollectionDB `json:",omitempty,omitzero"`
}
