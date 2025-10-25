package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type BannerDB struct {
	BannerOrder         int64
	StartDate           time.Time
	EndDate             time.Time
	Url                 string
	FileName            string
	WebViewTitle        string
	WebViewUrl          string
	LinkedLobbyBannerId int32
	BannerType          flatdata.EventContentType
	BannerDisplayType   BannerDisplayType
}
