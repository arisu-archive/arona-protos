package protos

import (
	"time"
)

type AccountBanByNexonDB struct {
	AccountId int64 `json:",omitempty,omitzero"`
	Npsn int64 `json:",omitempty,omitzero"`
	AccountBanId int64 `json:",omitempty,omitzero"`
	BanType int32 `json:",omitempty,omitzero"`
	BanDay int32 `json:",omitempty,omitzero"`
	BanStartDate time.Time `json:",omitempty,omitzero"`
	BanEndDate time.Time `json:",omitempty,omitzero"`
	BanComment string `json:",omitempty,omitzero"`
	DeleteFlag int32 `json:",omitempty,omitzero"`
}
