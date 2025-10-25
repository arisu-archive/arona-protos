package protos

type WeekDungeonBattleResultResponse struct {
	ResponsePacket
	Protocol                  Protocol
	WeekDungeonStageHistoryDB WeekDungeonStageHistoryDB
	LevelUpCharacterDBs       []CharacterDB
	ParcelResultDB            ParcelResultDB
}
