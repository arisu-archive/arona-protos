package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type AccountCurrencyDB struct {
	AccountLevel           int64
	AcademyLocationRankSum int64
	CurrencyDict           map[flatdata.CurrencyTypes]int64
	UpdateTimeDict         map[flatdata.CurrencyTypes]time.Time
}
