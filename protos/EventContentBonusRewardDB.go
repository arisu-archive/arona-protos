package protos

type EventContentBonusRewardDB struct {
	EventContentId     int64
	EventStageUniqueId int64
	BonusPercentage    int64
	BonusParcelInfo    ParcelInfo
}
