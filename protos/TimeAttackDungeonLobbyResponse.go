package protos

type TimeAttackDungeonLobbyResponse struct {
	ResponsePacket
	Protocol                Protocol
	RoomDBs                 map[int64]TimeAttackDungeonRoomDB
	PreviousRoomDB          TimeAttackDungeonRoomDB
	ParcelResultDB          ParcelResultDB
	AchieveSeasonBestRecord bool
	SeasonBestRecord        int64
}
