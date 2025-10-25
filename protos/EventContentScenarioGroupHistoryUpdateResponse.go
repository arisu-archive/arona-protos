package protos

type EventContentScenarioGroupHistoryUpdateResponse struct {
	ResponsePacket
	Protocol                  Protocol
	ScenarioGroupHistoryDBs   []ScenarioGroupHistoryDB
	EventContentCollectionDBs []EventContentCollectionDB
	ParcelResultDB            ParcelResultDB
}
