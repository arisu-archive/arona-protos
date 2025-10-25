package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type FurnitureDB struct {
	ConsumableItemBaseDB
	Type               flatdata.ParcelType
	Location           flatdata.FurnitureLocation
	CafeDBId           int64
	PositionX          float32
	PositionY          float32
	Rotation           float32
	ItemDeploySequence int64
}
