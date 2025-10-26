package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CraftNodeResult struct {
	NodeTier flatdata.CraftNodeTier `json:",omitempty,omitzero"`
	ParcelInfo ParcelInfo `json:",omitempty,omitzero"`
}
