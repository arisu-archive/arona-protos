package protos

type ScenarioPortalResponse struct {
	ResponsePacket
	Protocol                 Protocol
	StoryStrategyStageSaveDB StoryStrategyStageSaveDB
	ScenarioIds              []int64
}
