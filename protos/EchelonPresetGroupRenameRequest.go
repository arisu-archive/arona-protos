package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EchelonPresetGroupRenameRequest struct {
	RequestPacket
	Protocol         Protocol
	PresetGroupIndex int32
	ExtensionType    flatdata.EchelonExtensionType
	PresetGroupLabel string
}
