package protos

type StrategyClearRewardInfo struct {
	FirstClearReward        []ParcelInfo
	ThreeStarReward         []ParcelInfo
	StrategyObjectRewards   map[int64][]ParcelInfo
	ParcelResultDB          ParcelResultDB
	EventContentBonusReward []ParcelInfo
	CampaignStageHistoryDB  CampaignStageHistoryDB
}
