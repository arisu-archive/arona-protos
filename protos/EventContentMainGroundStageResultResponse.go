package protos

type EventContentMainGroundStageResultResponse struct {
	ResponsePacket
	Protocol                  Protocol
	TacticRank                int64
	CampaignStageHistoryDB    CampaignStageHistoryDB
	LevelUpCharacterDBs       []CharacterDB
	ParcelResultDB            ParcelResultDB
	FirstClearReward          []ParcelInfo
	ThreeStarReward           []ParcelInfo
	BonusReward               []ParcelInfo
	EventContentCollectionDBs []EventContentCollectionDB
}
