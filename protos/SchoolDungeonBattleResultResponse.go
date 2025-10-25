package protos

type SchoolDungeonBattleResultResponse struct {
	ResponsePacket
	Protocol                    Protocol
	SchoolDungeonStageHistoryDB SchoolDungeonStageHistoryDB
	LevelUpCharacterDBs         []CharacterDB
	FirstClearReward            []ParcelInfo
	ThreeStarReward             []ParcelInfo
	ParcelResultDB              ParcelResultDB
}
