package protos

type EventContentSubStageResultResponse struct {
	ResponsePacket
	Protocol                  Protocol
	TacticRank                int64
	CampaignStageHistoryDB    CampaignStageHistoryDB
	LevelUpCharacterDBs       []CharacterDB
	ParcelResultDB            ParcelResultDB
	FirstClearReward          []ParcelInfo
	BonusReward               []ParcelInfo
	EventContentCollectionDBs []EventContentCollectionDB
}
