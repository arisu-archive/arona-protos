package protos

import (
	"time"
)

type MomoTalkChoiceDB struct {
	CharacterDBId   int64
	MessageGroupId  int64
	ChosenMessageId int64
	ChosenDate      time.Time
}
