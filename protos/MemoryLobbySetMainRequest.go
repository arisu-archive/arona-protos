package protos

type MemoryLobbySetMainRequest struct {
	RequestPacket
	Protocol      Protocol
	MemoryLobbyId int64
}
