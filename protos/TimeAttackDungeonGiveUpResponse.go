package protos

type TimeAttackDungeonGiveUpResponse struct {
	ResponsePacket
	Protocol                Protocol
	RoomDB                  TimeAttackDungeonRoomDB
	ParcelResultDB          ParcelResultDB
	AchieveSeasonBestRecord bool
	SeasonBestRecord        int64
}
