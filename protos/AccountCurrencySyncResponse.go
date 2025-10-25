package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type AccountCurrencySyncResponse struct {
	ResponsePacket
	Protocol          Protocol
	AccountCurrencyDB AccountCurrencyDB
	ExpiredCurrency   map[flatdata.CurrencyTypes]int64
}
