package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ToastDB struct {
	UniqueId     int64
	Text         string
	LocalizeText map[flatdata.Language]string
	ToastId      string
	BeginDate    time.Time
	EndDate      time.Time
	LifeTime     int32
	Delay        int32
}
