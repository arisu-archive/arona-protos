package protos

import (
	"time"
)

type MomoTalkOutLineDB struct {
	CharacterDBId        int64
	CharacterId          int64
	LatestMessageGroupId int64
	ChosenMessageId      *int64
	LastUpdateDate       time.Time
}
