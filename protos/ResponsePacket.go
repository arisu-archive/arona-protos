package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ResponsePacket struct {
	BasePacket
	ServerTimeTicks                 int64
	ServerNotification              ServerNotificationFlag
	MissionProgressDBs              []MissionProgressDB
	EventMissionProgressDBDict      map[int64][]MissionProgressDB
	BattlePassMissionProgressDBDict map[int64][]MissionProgressDB
	StaticOpenConditions            map[flatdata.OpenConditionContent]OpenConditionLockReason
}
