package protos

type ScenarioListResponse struct {
	ResponsePacket
	Protocol                Protocol
	ScenarioHistoryDBs      []ScenarioHistoryDB
	ScenarioGroupHistoryDBs []ScenarioGroupHistoryDB
	ScenarioCollectionDBs   []ScenarioCollectionDB
}
