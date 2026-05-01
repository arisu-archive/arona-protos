package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type BannerDB struct {
	BannerOrder         int64                     `json:",omitempty,omitzero"`
	StartDate           MxTime                    `json:",omitempty,omitzero"`
	EndDate             MxTime                    `json:",omitempty,omitzero"`
	Url                 string                    `json:",omitempty,omitzero"`
	FileName            string                    `json:",omitempty,omitzero"`
	WebViewTitle        string                    `json:",omitempty,omitzero"`
	WebViewUrl          string                    `json:",omitempty,omitzero"`
	LinkedLobbyBannerId int32                     `json:",omitempty,omitzero"`
	BannerType          flatdata.EventContentType `json:",omitempty,omitzero"`
	BannerDisplayType   BannerDisplayType         `json:",omitempty,omitzero"`
}
