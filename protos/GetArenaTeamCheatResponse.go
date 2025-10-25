package protos

type GetArenaTeamCheatResponse struct {
	ResponsePacket
	Protocol Protocol
	Opponent ArenaUserDB
}
