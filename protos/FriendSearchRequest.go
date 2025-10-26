package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type FriendSearchRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	FriendCode string `json:",omitempty,omitzero"`
	LevelOption flatdata.FriendSearchLevelOption `json:",omitempty,omitzero"`
}
