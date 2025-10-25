package protos

type ScenarioSkipRequest struct {
	RequestPacket
	Protocol             Protocol
	ScriptGroupId        int64
	SkipPointScriptCount int32
}
