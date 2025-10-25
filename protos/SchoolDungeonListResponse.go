package protos

type SchoolDungeonListResponse struct {
	ResponsePacket
	Protocol                        Protocol
	SchoolDungeonStageHistoryDBList []SchoolDungeonStageHistoryDB
}
