package protos

type EchelonPresetGroupRenameResponse struct {
	ResponsePacket
	Protocol      Protocol
	PresetGroupDB EchelonPresetGroupDB
}
