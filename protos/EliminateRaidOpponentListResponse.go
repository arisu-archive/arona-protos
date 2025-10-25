package protos

type EliminateRaidOpponentListResponse struct {
	ResponsePacket
	Protocol        Protocol
	OpponentUserDBs []EliminateRaidUserDB
}
