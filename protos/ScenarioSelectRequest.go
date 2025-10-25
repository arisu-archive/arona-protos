package protos

type ScenarioSelectRequest struct {
	RequestPacket
	Protocol          Protocol
	ScriptGroupId     int64
	ScriptSelectGroup int64
}
