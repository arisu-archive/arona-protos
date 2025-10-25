package protos

type ScenarioConfirmMainStageResponse struct {
	ResponsePacket
	Protocol       Protocol
	ParcelResultDB ParcelResultDB
	SaveDataDB     StoryStrategyStageSaveDB
	ScenarioIds    []int64
}
