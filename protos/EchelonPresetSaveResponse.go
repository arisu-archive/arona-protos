package protos

type EchelonPresetSaveResponse struct {
	ResponsePacket
	Protocol Protocol
	PresetDB EchelonPresetDB
}
