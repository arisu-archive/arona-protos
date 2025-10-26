package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
	"time"
)

type BannerDB struct {
	BannerOrder int64 `json:",omitempty,omitzero"`
	StartDate time.Time `json:",omitempty,omitzero"`
	EndDate time.Time `json:",omitempty,omitzero"`
	Url string `json:",omitempty,omitzero"`
	FileName string `json:",omitempty,omitzero"`
	WebViewTitle string `json:",omitempty,omitzero"`
	WebViewUrl string `json:",omitempty,omitzero"`
	LinkedLobbyBannerId int32 `json:",omitempty,omitzero"`
	BannerType flatdata.EventContentType `json:",omitempty,omitzero"`
	BannerDisplayType BannerDisplayType `json:",omitempty,omitzero"`
}
