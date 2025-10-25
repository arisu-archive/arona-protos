package protos

type ClanLobbyResponse struct {
	ResponsePacket
	Protocol              Protocol
	IrcConfig             IrcServerConfig
	AccountClanDB         ClanDB
	DefaultExposedClanDBs []ClanDB
	AccountClanMemberDB   ClanMemberDB
	ClanMemberDBs         []ClanMemberDB
}
