package protos

type WeekDungeonListResponse struct {
	ResponsePacket
	Protocol                      Protocol
	AdditionalStageIdList         []int64
	WeekDungeonStageHistoryDBList []WeekDungeonStageHistoryDB
}
