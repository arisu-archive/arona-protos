package protos

type EchelonPresetSaveRequest struct {
	RequestPacket
	Protocol Protocol
	PresetDB EchelonPresetDB
}
