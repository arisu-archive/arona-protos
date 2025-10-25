package protos

type EventContentAdventureListResponse struct {
	ResponsePacket
	Protocol                   Protocol
	StageHistoryDBs            []CampaignStageHistoryDB
	StrategyObjecthistoryDBs   []StrategyObjectHistoryDB
	EventContentBonusRewardDBs []EventContentBonusRewardDB
	AlreadyReceiveRewardId     []int64
	StagePoint                 int64
}
