package protos

type EliminateRaidGetBestTeamResponse struct {
	ResponsePacket
	Protocol               Protocol
	RaidTeamSettingDBsDict map[string][]RaidTeamSettingDB
}
