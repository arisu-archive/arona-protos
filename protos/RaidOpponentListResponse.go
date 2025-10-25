package protos

type RaidOpponentListResponse struct {
	ResponsePacket
	Protocol        Protocol
	OpponentUserDBs []SingleRaidUserDB
}
