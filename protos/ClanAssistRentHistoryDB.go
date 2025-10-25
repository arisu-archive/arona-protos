package protos

import (
	"time"
)

type ClanAssistRentHistoryDB struct {
	AssistCharacterAccountId int64
	AssistCharacterDBId      int64
	RentDate                 time.Time
	AssistCharacterId        int64
}
