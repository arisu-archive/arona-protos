package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type FriendSearchRequest struct {
	RequestPacket
	Protocol    Protocol
	FriendCode  string
	LevelOption flatdata.FriendSearchLevelOption
}
