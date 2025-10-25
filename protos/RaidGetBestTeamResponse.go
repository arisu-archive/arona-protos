package protos

type RaidGetBestTeamResponse struct {
	ResponsePacket
	Protocol           Protocol
	RaidTeamSettingDBs []RaidTeamSettingDB
}
