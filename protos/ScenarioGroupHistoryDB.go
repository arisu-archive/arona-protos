package protos

import (
	"time"
)

type ScenarioGroupHistoryDB struct {
	AccountServerId int64 `json:",omitempty,omitzero"`
	ScenarioGroupUqniueId int64 `json:",omitempty,omitzero"`
	ScenarioType int64 `json:",omitempty,omitzero"`
	EventContentId *int64 `json:",omitempty,omitzero"`
	ClearDateTime time.Time `json:",omitempty,omitzero"`
	IsReturn bool `json:",omitempty,omitzero"`
}
