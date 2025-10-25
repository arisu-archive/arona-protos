package protos

type EventContentTacticResultResponse struct {
	ResponsePacket
	Protocol                  Protocol
	TacticRank                int64
	CampaignStageHistoryDB    CampaignStageHistoryDB
	LevelUpCharacterDBs       []CharacterDB
	FirstClearReward          []ParcelInfo
	StrategyObject            Strategy
	StrategyObjectRewards     map[int64][]ParcelInfo
	BonusReward               []ParcelInfo
	ParcelResultDB            ParcelResultDB
	SaveDataDB                EventContentMainStageSaveDB
	EventContentCollectionDBs []EventContentCollectionDB
}
