package protos

type ScenarioEndTurnResponse struct {
	ResponsePacket
	Protocol          Protocol
	SaveDataDB        StoryStrategyStageSaveDB
	AccountCurrencyDB AccountCurrencyDB
	ScenarioIds       []int64
}
