package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"time"
)

type ToastDB struct {
	UniqueId int64 `json:",omitempty,omitzero"`
	Text string `json:",omitempty,omitzero"`
	LocalizeText map[flatdata.Language]string `json:",omitempty,omitzero"`
	ToastId string `json:",omitempty,omitzero"`
	BeginDate time.Time `json:",omitempty,omitzero"`
	EndDate time.Time `json:",omitempty,omitzero"`
	LifeTime int32 `json:",omitempty,omitzero"`
	Delay int32 `json:",omitempty,omitzero"`
}
