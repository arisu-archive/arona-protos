package protos

type ScenarioMapMoveRequest struct {
	RequestPacket
	Protocol        Protocol
	StageUniqueId   int64
	EchelonEntityId int64
	DestPosition    HexLocation
}
