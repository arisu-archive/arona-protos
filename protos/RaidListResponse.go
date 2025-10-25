package protos

type RaidListResponse struct {
	ResponsePacket
	Protocol      Protocol
	CreateRaidDBs []RaidDB
	EnterRaidDBs  []RaidDB
	ListRaidDBs   []RaidDB
}
