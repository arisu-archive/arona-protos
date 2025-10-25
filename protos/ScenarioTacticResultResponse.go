package protos

type ScenarioTacticResultResponse struct {
	ResponsePacket
	Protocol       Protocol
	StrategyObject Strategy
	SaveDataDB     StoryStrategyStageSaveDB
	IsPlayerWin    bool
	ScenarioIds    []int64
}
