package protos

type WorldRaidLobbyRequest struct {
	RequestPacket
	Protocol Protocol
	SeasonId int64
}
