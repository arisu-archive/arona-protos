package protos

type ScenarioClearResponse struct {
	ResponsePacket
	Protocol              Protocol
	ScenarioHistoryDB     ScenarioHistoryDB
	ParcelResultDB        ParcelResultDB
	ScenarioCollectionDBs []ScenarioCollectionDB
}
