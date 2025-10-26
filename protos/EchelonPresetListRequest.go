package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EchelonPresetListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EchelonExtensionType flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
}
