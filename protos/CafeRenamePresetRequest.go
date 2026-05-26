package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CafeRenamePresetRequest struct {
	RequestPacket
	PresetType flatdata.CafePresetType `json:",omitempty,omitzero"`
	SlotId     int32                   `json:",omitempty,omitzero"`
	PresetName string
}
