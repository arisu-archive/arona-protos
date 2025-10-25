package protos

type ScenarioEnterTacticRequest struct {
	RequestPacket
	Protocol      Protocol
	StageUniqueId int64
	EchelonIndex  int64
	EnemyIndex    int64
}
