package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type RaidListRequest struct {
	RequestPacket
	Protocol           Protocol
	RaidBossGroup      string
	RaidDifficulty     flatdata.Difficulty
	RaidRoomSortOption RaidRoomSortOption
}
