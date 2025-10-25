package protos

type ShopGachaRecruitListResponse struct {
	ResponsePacket
	Protocol                  Protocol
	ShopRecruits              []ShopRecruitDB
	ShopFreeRecruitHistoryDBs []ShopFreeRecruitHistoryDB
}
