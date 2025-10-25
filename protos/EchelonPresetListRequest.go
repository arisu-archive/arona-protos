package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EchelonPresetListRequest struct {
	RequestPacket
	Protocol             Protocol
	EchelonExtensionType flatdata.EchelonExtensionType
}
