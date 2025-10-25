package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ParcelCost struct {
	ParcelInfos      []ParcelInfo
	Currency         CurrencyTransaction
	EquipmentDBs     []EquipmentDB
	ItemDBs          []ItemDB
	FurnitureDBs     []FurnitureDB
	ConsumeCondition flatdata.ConsumeCondition
}
