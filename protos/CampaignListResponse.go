package protos

type CampaignListResponse struct {
	ResponsePacket
	Protocol                             Protocol
	CampaignChapterClearRewardHistoryDBs []CampaignChapterClearRewardHistoryDB
	StageHistoryDBs                      []CampaignStageHistoryDB
	StrategyObjecthistoryDBs             []StrategyObjectHistoryDB
}
