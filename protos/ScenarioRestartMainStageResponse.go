package protos

type ScenarioRestartMainStageResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	SaveDataDB     StoryStrategyStageSaveDB
}
