package protos

type RaidCompleteListResponse struct {
	ResponsePacket
	Protocol          Protocol
	RaidDBs           []RaidDB
	StackedDamage     int64
	ReceiveRewardId   []int64
	CurSeasonUniqueId int64
}
