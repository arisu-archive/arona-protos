package protos

import (
	"time"
)

type ResetableContentValueDB struct {
	ResetableContentId ResetableContentId
	ContentValue       int64
	LastUpdateTime     time.Time
}
