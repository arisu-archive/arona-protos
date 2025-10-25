package protos

type MemoryLobbyListResponse struct {
	ResponsePacket
	Protocol       Protocol
	MemoryLobbyDBs []MemoryLobbyDB
}
