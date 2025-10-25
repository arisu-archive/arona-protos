package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type AccountDB struct {
	ServerId                   int64
	Nickname                   string
	CallName                   string
	CallNameKatakana           string
	CallNameKorean             string
	DevId                      string
	State                      flatdata.AccountState
	Level                      int32
	Exp                        int64
	Comment                    string
	LobbyMode                  int32
	RepresentCharacterServerId int64
	MemoryLobbyUniqueId        int64
	LastConnectTime            time.Time
	BirthDay                   time.Time
	CallNameUpdateTime         time.Time
	PublisherAccountId         int64
	RetentionDays              *int32
	VIPLevel                   *int32
	CreateDate                 time.Time
	UnReadMailCount            *int32
	LinkRewardDate             *time.Time
}
