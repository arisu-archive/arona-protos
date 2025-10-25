package protos

type CampaignTacticResultResponse struct {
	ResponsePacket
	Protocol               Protocol
	TacticRank             int64
	CampaignStageHistoryDB CampaignStageHistoryDB
	LevelUpCharacterDBs    []CharacterDB
	FirstClearReward       []ParcelInfo
	ThreeStarReward        []ParcelInfo
	StrategyObject         Strategy
	StrategyObjectRewards  map[int64][]ParcelInfo
	ParcelResultDB         ParcelResultDB
	SaveDataDB             CampaignMainStageSaveDB
}
