package protos

type ScenarioMapMoveResponse struct {
	ResponsePacket
	Protocol                  Protocol
	SaveDataDB                StoryStrategyStageSaveDB
	ScenarioIds               []int64
	EchelonEntityId           int64
	StrategyObject            Strategy
	StrategyObjectParcelInfos []ParcelInfo
}
