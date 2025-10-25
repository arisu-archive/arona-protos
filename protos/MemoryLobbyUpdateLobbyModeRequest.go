package protos

type MemoryLobbyUpdateLobbyModeRequest struct {
	RequestPacket
	Protocol          Protocol
	IsMemoryLobbyMode bool
}
