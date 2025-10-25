package protos

type RaidGetBestTeamRequest struct {
	RequestPacket
	Protocol        Protocol
	SearchAccountId int64
}
