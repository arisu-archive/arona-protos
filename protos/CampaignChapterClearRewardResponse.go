package protos

type CampaignChapterClearRewardResponse struct {
	ResponsePacket
	Protocol                            Protocol
	CampaignChapterClearRewardHistoryDB CampaignChapterClearRewardHistoryDB
	ParcelResultDB                      ParcelResultDB
}
