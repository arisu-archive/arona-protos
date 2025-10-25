package protos

type EchelonPresetListResponse struct {
	ResponsePacket
	Protocol       Protocol
	PresetGroupDBs []EchelonPresetGroupDB
}
