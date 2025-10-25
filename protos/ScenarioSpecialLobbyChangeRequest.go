package protos

type ScenarioSpecialLobbyChangeRequest struct {
	RequestPacket
	Protocol            Protocol
	MemoryLobbyId       int64
	MemoryLobbyIdBefore int64
}
