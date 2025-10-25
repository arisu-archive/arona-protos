package protos

type EventContentEnterTacticRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	StageUniqueId  int64
	EchelonIndex   int64
	EnemyIndex     int64
}
